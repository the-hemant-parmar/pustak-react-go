package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Addr string `mapstructure:"addr"`
	}
	Mongo struct {
		URI      string `mapstructure:"uri"`
		Database string `mapstructure:"database"`
	}
}

func Load(path string) *Config {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
