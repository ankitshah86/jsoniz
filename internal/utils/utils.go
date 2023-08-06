package utils

import (
	"fmt"
	"os"
	"strconv"
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
		//try to convert to int
		val, err := strconv.Atoi(osValue)
		if err != nil {
			fmt.Println("Could not convert to int")
			return defaultValue
		} else {
			return val
		}
	}

	return defaultValue
}
