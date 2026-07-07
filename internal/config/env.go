package config

import (
	"os"
	"strconv"
)

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

func getIntEnv(key string, fallback int) int {
	value := os.Getenv(key)

	result, err := strconv.Atoi(value)

	if err != nil {
		return fallback
	}

	return result
}

func getFloatEnv(key string, fallback float64) float64 {
	value := os.Getenv(key)

	result, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return fallback
	}

	return result
}

func getBoolEnv(key string, fallback bool) bool {
	value := os.Getenv(key)

	result, err := strconv.ParseBool(value)

	if err != nil {
		return fallback
	}

	return result
}
