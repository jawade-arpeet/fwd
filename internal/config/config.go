package config

import (
	"fwd/internal/constants"

	"github.com/go-playground/validator/v10"
)

type ServerConfig struct {
	Port   string        `validate:"required,numeric"`
	RunEnv constants.Env `validate:"required,oneof=dev stg prod"`
}

type Config struct {
	Server *ServerConfig `validate:"required"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Server: &ServerConfig{
			Port:   "8080",
			RunEnv: constants.Dev,
		},
	}

	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
