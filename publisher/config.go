package main

import (
	"os"
)

type Config struct {
	HTTPPort string
}

func LoadConfig() Config {
	return Config{
		HTTPPort: GetEnv("HTTP_PORT", "3000"),
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
