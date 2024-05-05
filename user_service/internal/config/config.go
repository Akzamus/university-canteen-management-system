package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	Database DatabaseConfig
	Server   ServerConfig
}

func NewAppConfig() (AppConfig, error) {
	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return AppConfig{}, err
	}

	httpPortStr := os.Getenv("HTTP_PORT")
	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		return AppConfig{}, err
	}

	httpTimeoutSecondsStr := os.Getenv("HTTP_TIMEOUT_SECONDS")
	httpTimeoutSeconds, err := strconv.Atoi(httpTimeoutSecondsStr)
	if err != nil {
		return AppConfig{}, err
	}

	dbConfig := DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     dbPort,
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
	}

	serverConfig := ServerConfig{
		HttpPort:           httpPort,
		HttpTimeoutSeconds: httpTimeoutSeconds,
	}

	return AppConfig{
		Database: dbConfig,
		Server:   serverConfig,
	}, nil
}
