package config

import (
    "log"
    "os"
)

type Configuration struct {
    ApiUrl string
}

var Config Configuration

func LoadConfig() {
    Config = Configuration{
        ApiUrl: getEnv("EXTERNAL_API_URL", "https://api.restful-api.dev/objects"),
    }
}

func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        log.Printf("Using default value for %s: %s\n", key, defaultValue)
        return defaultValue
    }
    return value
}