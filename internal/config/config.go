package config

import (
	"fwd/internal/constants"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ServerConfig struct {
	Port   string        `validate:"required,numeric" mapstructure:"port"`
	RunEnv constants.Env `validate:"required,oneof=dev stg prod" mapstructure:"run_env"`
}

type Config struct {
	Server *ServerConfig `validate:"required" mapstructure:"server"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		zap.L().Error("error while reading config file", zap.String("operation", "config.LoadConfig"), zap.Error(err))
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		zap.L().Error(
			"error while unmarshaling config to struct",
			zap.String("operation", "config.LoadConfig"),
			zap.Error(err),
		)
		return nil, err
	}

	if err := validator.New().Struct(cfg); err != nil {
		zap.L().Error(
			"invalid config struct",
			zap.String("operation", "config.LoadConfig"),
			zap.Error(err),
		)
		return nil, err
	}

	return &cfg, nil
}
