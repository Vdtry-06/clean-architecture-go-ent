package config

import (
	"github.com/spf13/viper"
)

type Config struct {
    Port string
    DB   string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("port", "8080")
	viper.SetDefault("db", "sqlite://file:ent?mode=memory")	

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		Port: viper.GetString("port"),
		DB:   viper.GetString("db"),
	}, nil
}