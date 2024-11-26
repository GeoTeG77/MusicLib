package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", fmt.Errorf("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	return connectionString, nil
}
