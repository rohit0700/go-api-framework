package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		IP   string
		Port string
	}
	Database struct {
		Driver   string
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

var cfg *Config

func InitConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
