package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Token string `json:"token"`
}

func LoadConfig() Config {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		panic(err.Error())
	}
	return config
}
