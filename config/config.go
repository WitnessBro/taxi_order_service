package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Address string
}

func NewConfig() (*Config, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("config/")
	if err := vp.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("viper can’t read config: %w", err)
	}
	vp.SetDefault("address", "localhost:8080")
	config := Config{}
	if err := vp.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("viper can't unmarshal config: %w", err)
	}
	return &config, nil
}
