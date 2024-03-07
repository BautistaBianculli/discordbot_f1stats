package main

import (
	"BotDiscordGO/internal/server/infra"
	"BotDiscordGO/internal/server/infra/config"
)

func main() {
	c := config.LoadConfig()

	infra.InitApp(c)

}
