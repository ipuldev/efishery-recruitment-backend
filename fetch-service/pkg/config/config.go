package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Name string
	Type string
	Path string
}

func (config *Config) Init() error {
	viper.SetConfigName(config.Name)
	viper.SetConfigType(config.Type)
	viper.AddConfigPath(config.Path)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
