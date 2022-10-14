package util

import (
	"os"
	"strconv"
	"strings"
)

func EnvStr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func EnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if num, err := strconv.Atoi(value); err == nil {
			return num
		}
	}

	return fallback
}

func EnvBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if b, err := strconv.ParseBool(value); err == nil {
			return b
		}
	}

	return fallback
}

func EnvStringSlice(key, split string, fallback []string) []string {
	if value, ok := os.LookupEnv(key); ok {
		return strings.Split(value, split)
	}

	return fallback
}
