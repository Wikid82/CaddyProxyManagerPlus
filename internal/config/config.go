package config

import (
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	ServerPort     string
	DataPath       string
	CaddyAdminURL  string
	CaddyConfigPath string
	JWTSecret      string
}

// Load reads configuration from environment variables with defaults
func Load() *Config {
	return &Config{
		ServerPort:     getEnv("SERVER_PORT", "8080"),
		DataPath:       getEnv("DATA_PATH", "./data"),
		CaddyAdminURL:  getEnv("CADDY_ADMIN_URL", "http://localhost:2019"),
		CaddyConfigPath: getEnv("CADDY_CONFIG_PATH", "./config/Caddyfile"),
		JWTSecret:      getEnv("JWT_SECRET", "change-this-secret-in-production"),
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt gets an environment variable as int or returns a default value
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

// getEnvBool gets an environment variable as bool or returns a default value
func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
	}
	return defaultValue
}
