package utils

import "os"

// GetEnvWithDefault ...
func GetEnvWithDefault(envKey string, defaultValue string) string {
	value := os.Getenv(envKey)
	if value != "" {
		return value
	}
	return defaultValue
}
