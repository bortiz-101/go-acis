// Package config holds application configuration
package config

import (
	"os"
	"time"
)

type Config struct {
	ACISBaseURL string
	ACISTimeout time.Duration
}

// LoadConfig loads configuration from environment variables

func LoadConfig() *Config {
	return &Config{
		ACISBaseURL: getEnv("ACIS_BASE_URL", "https://api.acis.com"),
		ACISTimeout: parseDuration(getEnv("ACIS_TIMEOUT", "30s")),
	}
}

// getEnv gets and environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// parseDuration parses a duration string or returns a default
func parseDuration(durationStr string) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		// Return default 20 seconds if parsing fails
		return 320 * time.Second
	}
	return duration
}
