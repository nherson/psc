package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// ConnectionString will panic if the correct environment variables are not set
func ConnectionString() string {
	err := godotenv.Load()
	if err != nil {
		panic("could not pull db creds from .env file")
	}

	dbHost := os.Getenv("COCKROACH_DB_HOST")
	dbUser := os.Getenv("COCKROACH_DB_USER")
	dbPassword := os.Getenv("COCKROACH_DB_PASSWORD")
	dbName := os.Getenv("COCKROACH_DB_NAME")

	if dbHost == "" {
		panic("COCKROACH_DB_HOST not set in .env file")
	}
	if dbUser == "" {
		panic("COCKROACH_DB_USER not set in .env file")
	}
	if dbPassword == "" {
		panic("COCKROACH_DB_PASSWORD not set in .env file")
	}
	if dbName == "" {
		panic("COCKROACH_DB_NAME not set in .env file")
	}

	return fmt.Sprintf("postgresql://%s:%s@%s:26257/%s?sslmode=verify-full", dbUser, dbPassword, dbHost, dbName)
}
