package utils

import (
	"os"
	"strings"
)

// GetEnvBool returns the boolean value of the environment variable
func GetEnvBool(key string, defaultValue bool) bool {
	osValue := os.Getenv(key)
	if strings.ToLower(osValue) == "true" {
		return true
	} else if strings.ToLower(osValue) == "false" {
		return false
	}

	return defaultValue
}

// GetEnvString returns the string value of the environment variable
func GetEnvString(key string, defaultValue string) string {
	osValue := os.Getenv(key)
	if osValue != "" {
		return osValue
	}

	return defaultValue
}

// GetEnvInt returns the integer value of the environment variable
func GetEnvInt(key string, defaultValue int) int {
	osValue := os.Getenv(key)
	if osValue != "" {
		return defaultValue
	}

	return defaultValue
}
