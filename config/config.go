package config

import (
	"fmt"
	"os"
	"strconv"
)

// Server contains the HTTP server configurations.
type ServerConfigs struct {
	Port             string
	GracefulShutdown bool
	ServiceName      string
}

// NewServer creates a new Server instance with default values.
func GetServerConfigs() *ServerConfigs {
	return &ServerConfigs{
		Port:             fmt.Sprintf(":%s", getEnvOrDefault("PORT", "8080")),
		GracefulShutdown: getEnvOrDefaultBool("ENABLE_GRACEFUL_SHUTDOWN", true),
		ServiceName:      getEnvOrDefault("SERVICE_NAME", "go-app"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvOrDefaultBool(key string, defaultValue bool) bool {
	value, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return defaultValue
	}
	return value
}
