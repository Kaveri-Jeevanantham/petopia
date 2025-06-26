package config

import (
	"os"
)

type Config struct {
	// Server config
	Port string
	Host string

	// Database config
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string

	// RabbitMQ config
	RabbitMQHost     string
	RabbitMQPort     string
	RabbitMQUser     string
	RabbitMQPassword string
	RabbitMQVhost    string

	// CORS config
	CORSAllowedOrigin string
}

func Load() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		Host: getEnv("HOST", "localhost"),

		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBHost:     getEnv("DB_HOST", ""),
		DBName:     getEnv("DB_NAME", ""),

		RabbitMQHost:     getEnv("RABBITMQ_HOST", ""),
		RabbitMQPort:     getEnv("RABBITMQ_PORT", ""),
		RabbitMQUser:     getEnv("RABBITMQ_USER", ""),
		RabbitMQPassword: getEnv("RABBITMQ_PASSWORD", ""),
		RabbitMQVhost:    getEnv("RABBITMQ_VHOST", ""),

		CORSAllowedOrigin: getEnv("CORS_ALLOWED_ORIGIN", "*"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
