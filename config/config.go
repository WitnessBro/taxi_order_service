package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Hostname string
	Port     string
	BasePath string
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("config/")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}
