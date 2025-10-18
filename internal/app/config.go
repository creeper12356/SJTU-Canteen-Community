package app

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `yaml:"app" mapstructure:"app"`
	Database DatabaseConfig `yaml:"database" mapstructure:"database"`
	Redis    RedisConfig    `yaml:"redis" mapstructure:"redis"`
}

type AppConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver" mapstructure:"driver"`
	Source string `yaml:"source" mapstructure:"source"`
}

type RedisConfig struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}

var Conf *Config

func LoadConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	Conf = &Config{}
	err = viper.Unmarshal(Conf)
	if err != nil {
		return fmt.Errorf("error unmarshaling config: %w", err)
	}

	return nil
}
