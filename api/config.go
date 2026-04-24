package main

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port                     int
	DatabaseConnectionString string
}

func NewConfig() (*Config, error) {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		return nil, errors.New("PORT not set")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.New("PORT is invalid")
	}

	dbConnectionString := os.Getenv("DATABASE_CONNECTION_STRING")
	if dbConnectionString == "" {
		return nil, errors.New("DATABASE_CONNECTION_STRING not set")
	}

	return &Config{
		Port:                     port,
		DatabaseConnectionString: dbConnectionString,
	}, nil
}
