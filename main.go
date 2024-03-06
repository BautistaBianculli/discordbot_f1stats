package main

import (
	"BotDiscordGO/internal/application/infra/repo"
	"BotDiscordGO/internal/server/infra/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := config.LoadConfig()
	// Token de tu bot, obtén esto al crear tu bot en el portal de desarrolladores de Discord

	// Crea una nueva sesión de DiscordGo
	dg, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		fmt.Println("Error al crear la sesión de DiscordGo,", err)
		return
	}

	// Registra un evento para ser llamado cada vez que se recibe un mensaje
	dg.AddHandler(repo.MessageCreate)

	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	// Abre la conexión al servidor de Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Error al abrir la conexión de Discord,", err)
		return
	}
	defer func(dg *discordgo.Session) {
		_ = dg.Close()
	}(dg)

	fmt.Println("Bot está ahora en funcionamiento. Presiona Ctrl+C para cerrar.")

	// Espera hasta que se reciba una señal para cerrar el bot
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
