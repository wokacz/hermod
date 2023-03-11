package env

import (
	"log"
	"os"
)

// Get - get string value from environment
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
