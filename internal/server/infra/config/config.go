package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Token          string `json:"token" required:"true"`
	AppId          string `json:"app_id" split_words:"true" required:"true"`
	GuildMainId    string `json:"guild_main_id" split_words:"true" required:"true"`
	DriverTableUrl string `json:"driver_table_url" split_words:"true"`
}

func LoadConfig() Config {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		panic(err.Error())
	}
	return config
}
