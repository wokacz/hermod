package env

import (
	"log"
	"os"
	"strconv"
)

// Get function is used to get an environment variable.
// If the environment variable is not found, the function will log a fatal error.
// If a default value is provided, the function will return the default value if the environment variable is not found.
func Get(key string, defaultValue ...string) (value string) {
	value = os.Getenv(key)
	if value == "" && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if value == "" {
		log.Fatalln("Env key \"" + key + "\" not found")
	}
	return value
}

// GetBool function is used to get a boolean environment variable.
// If the environment variable is not found, the function will log a fatal error.
// If a default value is provided, the function will return the default value if the environment variable is not found.
func GetBool(key string, defaultValue ...bool) (value bool) {
	value = os.Getenv(key) == "true"
	if !value && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !value {
		log.Fatalln("Env key \"" + key + "\" not found")
	}
	return value
}

// GetInt function is used to get an integer environment variable.
// If the environment variable is not found, the function will log a fatal error.
// If a default value is provided, the function will return the default value if the environment variable is not found.
func GetInt(key string, defaultValue ...int) (value int) {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if err != nil {
		log.Fatalln("Env key \"" + key + "\" not found")
	}
	return value
}

// GetInt64 function is used to get an integer environment variable.
// If the environment variable is not found, the function will log a fatal error.
// If a default value is provided, the function will return the default value if the environment variable is not found.
func GetInt64(key string, defaultValue ...int64) (value int64) {
	value, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if err != nil {
		log.Fatalln("Env key \"" + key + "\" not found")
	}
	return value
}

// GetUint64 function is used to get an integer environment variable.
// If the environment variable is not found, the function will log a fatal error.
// If a default value is provided, the function will return the default value if the environment variable is not found.
func GetUint64(key string, defaultValue ...uint64) (value uint64) {
	value, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if err != nil {
		log.Fatalln("Env key \"" + key + "\" not found")
	}
	return value
}

// GetFloat64 function is used to get a float environment variable.
// If the environment variable is not found, the function will log a fatal error.
// If a default value is provided, the function will return the default value if the environment variable is not found.
func GetFloat64(key string, defaultValue ...float64) (value float64) {
	// ParseFloat will return an error if the value is not a float.
	value, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if err != nil {
		log.Fatalln("Env key \"" + key + "\" not found")
	}
	return value
}
