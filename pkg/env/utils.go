package env

import (
	"os"
	"strconv"
	"strings"
)

func GetEnvOrStr(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func GetEnvOrInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return fallback
	}

	if v, err := strconv.Atoi(val); err == nil {
		return v
	}

	return fallback
}

func GetEnvOrBool(key string, fallback bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return fallback
	}

	return strings.ToLower(val) == "true"
}
