package config

import "os"

type Config struct {
	AppPort string
	DBUrl   string
	Env     string
}

func Load() *Config {
	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		DBUrl:   getEnv(
			"DATABASE_URL",
			"postgres://postgres:postgres@localhost:5432/userdb?sslmode=disable",
		),
		Env: getEnv("APP_ENV", "development"),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
