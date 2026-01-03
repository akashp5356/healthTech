package config

import (
	"os"
)

type Config struct {
	AppPort    string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	RedisHost  string
	RedisPort  string
	UploadDir  string
	JWTSecret  string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func LoadConfig() *Config {
	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBName:     getEnv("DB_NAME", "healthtech"),
		RedisHost:  getEnv("REDIS_HOST", "localhost"),
		RedisPort:  getEnv("REDIS_PORT", "6379"),
		UploadDir:  getEnv("UPLOAD_DIR", "./uploads"),
		JWTSecret:  getEnv("JWT_SECRET", "supersecretkey"),
	}
}
