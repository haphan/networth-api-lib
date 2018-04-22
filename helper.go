package main

import (
	"os"
)

// GetEnv get env var with default fallback
func GetEnv(key, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultVal
}
