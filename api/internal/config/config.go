package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment string

const (
	EnvDevelopment Environment = "development"
	EnvProduction  Environment = "production"
)

type Config struct {
	Domain                   string
	Port                     int
	DatabaseConnectionString string
	Environment              Environment
	JWTSecret                string
	ResendKey                string
}

func New() (*Config, []error) {
	_ = godotenv.Load()
	issues := []error{}

	env, err := getEnvironment()
	if err != nil {
		issues = append(issues, err)
	}

	port, err := getPort()
	if err != nil {
		issues = append(issues, err)
	}

	dbConnString, err := getDatabaseConnectionString()
	if err != nil {
		issues = append(issues, err)
	}

	domain, err := getDomain()
	if err != nil {
		issues = append(issues, err)
	}

	resendKey, err := getResendApiKey()
	if err != nil {
		issues = append(issues, err)
	}

	jwtSecret, err := getJWTSecret()
	if err != nil {
		issues = append(issues, err)
	}

	return &Config{
		Port:                     port,
		DatabaseConnectionString: dbConnString,
		Environment:              env,
		Domain:                   domain,
		JWTSecret:                jwtSecret,
		ResendKey:                resendKey,
	}, issues
}

func getEnvironment() (Environment, error) {
	envStr := os.Getenv("ENV")
	if envStr == "" {
		envStr = "development"
	}

	env := Environment(envStr)

	switch env {
	case EnvDevelopment, EnvProduction:
		return env, nil
	default:
		return "", errors.New("ENV must be development|production")
	}
}

func getPort() (int, error) {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		return 0, errors.New("PORT not set")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, errors.New("PORT is invalid")
	}
	return port, nil
}

func getDatabaseConnectionString() (string, error) {
	connString := os.Getenv("DATABASE_CONNECTION_STRING")
	if connString == "" {
		return "", errors.New("DATABASE_CONNECTION_STRING not set")
	}
	return connString, nil
}

func getDomain() (string, error) {
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		return "", errors.New("DOMAIN not set")
	}
	return domain, nil
}

func getResendApiKey() (string, error) {
	domain := os.Getenv("RESEND_API_KEY")
	if domain == "" {
		return "", errors.New("RESEND_API_KEY not set")
	}
	return domain, nil
}

func getJWTSecret() (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET not set")
	}

	if len(secret) < 32 {
		return "", errors.New("JWT_SECRET must be at least 32 characters")
	}
	return secret, nil
}
