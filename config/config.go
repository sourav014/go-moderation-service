package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	DB     DatabaseConfig
}

type ServerConfig struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	Debug        bool          `env:"SERVER_DEBUG,required"`
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DBName   string `env:"DB_NAME,required"`
}

func getEnvAsInt(key string, defaultValue int) int {
	valStr := os.Getenv(key)
	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Printf("Warning: %s is not set or invalid, using default: %d", key, defaultValue)
		return defaultValue
	}
	return val
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valStr := os.Getenv(key)
	val, err := time.ParseDuration(valStr)
	if err != nil {
		log.Printf("Warning: %s is not set or invalid, using default: %s", key, defaultValue)
		return defaultValue
	}
	return val
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valStr := os.Getenv(key)
	val, err := strconv.ParseBool(valStr)
	if err != nil {
		log.Printf("Warning: %s is not set or invalid, using default: %t", key, defaultValue)
		return defaultValue
	}
	return val
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverConfig := ServerConfig{
		Port:  getEnvAsInt("SERVER_PORT", 8080),
		Debug: getEnvAsBool("SERVER_DEBUG", false),
	}

	databaseConfig := DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     getEnvAsInt("DB_PORT", 5432),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
	}

	return &Config{
		Server: serverConfig,
		DB:     databaseConfig,
	}
}
