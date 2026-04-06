package config

import "fwd/internal/constants"

type ServerConfig struct {
	Port   string        `validate:"required,numeric"`
	RunEnv constants.Env `validate:"required,oneof=dev stg prod"`
}

type Config struct {
	Server *ServerConfig `validate:"required"`
}

func LoadConfig() (*Config, error) {
	return &Config{
		Server: &ServerConfig{
			Port:   "8080",
			RunEnv: constants.Dev,
		},
	}, nil
}
