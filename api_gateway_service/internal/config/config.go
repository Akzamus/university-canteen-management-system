package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	Server       ServerConfig
	UserService  UserServiceConfig
	JwtSecretKey string
}

func NewAppConfig() (AppConfig, error) {
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

	userServicePortStr := os.Getenv("USER_SERVICE_PORT")
	userServicePort, err := strconv.Atoi(userServicePortStr)
	if err != nil {
		return AppConfig{}, err
	}

	userServiceConfig := UserServiceConfig{
		Protocol: os.Getenv("USER_SERVICE_PROTOCOL"),
		Host:     os.Getenv("USER_SERVICE_HOST"),
		Port:     userServicePort,
	}

	serverConfig := ServerConfig{
		HttpPort:           httpPort,
		HttpTimeoutSeconds: httpTimeoutSeconds,
	}

	return AppConfig{
		Server:       serverConfig,
		UserService:  userServiceConfig,
		JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}, nil
}
