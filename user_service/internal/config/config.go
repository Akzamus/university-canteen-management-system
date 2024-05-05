package config

import (
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	Database DatabaseConfig
	Server   ServerConfig
}

func LoadConfigFromEnv(path string) (AppConfig, error) {
	err := godotenv.Load(path)
	if err != nil {
		return AppConfig{}, err
	}

	dbConfig := DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
	}

	serverConfig := ServerConfig{
		HttpPort: os.Getenv("HTTP_PORT"),
	}

	return AppConfig{
		Database: dbConfig,
		Server:   serverConfig,
	}, nil
}
