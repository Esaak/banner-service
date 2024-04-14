package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config represents the application configuration
type Config struct {
	Port        int
	PostgresURL string
	UserSecret  string
	AdminSecret string
}

// LoadConfig loads the configuration from environment variables or .env file
func LoadConfig() (Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return Config{}, err
	}
	psqInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), port, os.Getenv("DB_NAME"))
	config := Config{
		Port:        port,
		PostgresURL: psqInfo,
		UserSecret:  os.Getenv("USER_SECRET"),
		AdminSecret: os.Getenv("ADMIN_SECRET"),
	}

	return config, nil
}
