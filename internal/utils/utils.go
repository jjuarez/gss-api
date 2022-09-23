package utils

import "os"

// Getenv ...
func Getenv(envKey string, defaultValue string) string {
	envValue := os.Getenv(envKey)
	if "" == envValue {
		envValue = defaultValue
	}
	return envValue
}
