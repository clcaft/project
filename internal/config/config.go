package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Database DatabaseConfig
	APIPort  string
	LogLevel string
}

type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnvInt("DB_PORT", 3306),
			User:            getEnv("DB_USER", "root"),
			Password:        getEnv("DB_PASSWORD", ""),
			Name:            getEnv("DB_NAME", "procurement_db"),
			MaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 10),
			ConnMaxLifetime: getEnvDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
			ConnMaxIdleTime: getEnvDuration("DB_CONN_MAX_IDLE_TIME", 2*time.Minute),
		},
		APIPort:  getEnv("API_PORT", "8080"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
}

func (c *DatabaseConfig) DSN() string {
	return c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + strconv.Itoa(c.Port) + ")/" + c.Name + "?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci"
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if d, err := time.ParseDuration(value); err == nil {
			return d
		}
	}
	return defaultValue
}
