package helpers

import (
	"os"
)

func GetEnvVariable(key string) string {
	env := os.Getenv(key)
	if env == "" {
		panic("Environment variable " + key + " not found")
	}

	return env
}
