package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                     int
	DatabaseConnectionString string
}

func NewConfig() (*Config, error) {
	godotenv.Load()
	issues := []error{}

	portStr := os.Getenv("PORT")

	port, err := strconv.Atoi(portStr)
	slog.Debug(fmt.Sprintf("PORT: %v", port))
	if portStr == "" || err != nil {
		issues = append(issues, errors.New("PORT is invalid"))
	}

	dbConnectionString := os.Getenv("DATABASE_CONNECTION_STRING")
	slog.Debug(fmt.Sprintf("DATABASE_CONNECTION_STRING: %v", dbConnectionString))
	if dbConnectionString == "" {
		issues = append(issues, errors.New("DATABASE_CONNECTION_STRING not set"))
	}

	if len(issues) > 0 {
		for _, issue := range issues {
			slog.Error(issue.Error())
		}

		return nil, errors.New("invalid environment variables")
	}

	return &Config{
		Port:                     port,
		DatabaseConnectionString: dbConnectionString,
	}, nil
}
