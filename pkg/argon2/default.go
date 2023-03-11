package argon2

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/wokacz/hermod/pkg/env"
	"strconv"

	"golang.org/x/crypto/argon2"
)

var instance *Config

func Init() {
	memory, err := strconv.ParseUint(env.Get("ARGON2_MEMORY", "65536"), 10, 64)
	if err != nil {
		return
	}

	iterations, err := strconv.ParseInt(env.Get("ARGON2_ITERATIONS", "3"), 10, 0)
	if err != nil {
		return
	}

	keyLength, err := strconv.ParseInt(env.Get("ARGON2_KEY_LENGTH", "2"), 10, 0)
	if err != nil {
		return
	}

	saltLength, err := strconv.ParseInt(env.Get("ARGON2_SALT_LENGTH", "16"), 10, 0)
	if err != nil {
		return
	}

	parallelism, err := strconv.ParseInt(env.Get("ARGON2_PARALLELISM", "32"), 10, 0)
	if err != nil {
		return
	}

	instance = &Config{
		memory:      uint32(memory),
		iterations:  uint32(iterations),
		parallelism: uint8(parallelism),
		saltLength:  uint32(saltLength),
		keyLength:   uint32(keyLength),
	}
}

func Hash(password string) (encodedHash string, err error) {
	salt, err := generateRandomBytes(instance.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, instance.iterations, instance.memory, instance.parallelism, instance.keyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, instance.memory, instance.iterations, instance.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func Compare(password, encodedHash string) (match bool, err error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}

	return false, nil
}
