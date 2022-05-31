package env

import (
	"os"
	"strconv"
)

func GetEnvOrStr(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetEnvOrInt(key, fallback string) int {
	val, err := strconv.Atoi(GetEnvOrStr(key, fallback))
	if err != nil {
		return 0
	}
	return val
}
