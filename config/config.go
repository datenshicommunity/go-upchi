package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the .env file
func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

// GetEnv retrieves the value of the environment variable key
func GetEnv(key string) string {
    value := os.Getenv(key)
    if value == "" {
        log.Fatalf("Environment variable %s not set", key)
    }
    return value
}
