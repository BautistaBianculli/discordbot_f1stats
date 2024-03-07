package infra

import (
	"BotDiscordGO/internal/server/infra/config"
	"BotDiscordGO/internal/server/infra/config/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func InitApp(c config.Config) {

	dg, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		fmt.Println("Error al crear la sesión de DiscordGo,", err)
		return
	}

	// Registra un evento para ser llamado cada vez que se recibe un mensaje
	dg.AddHandler(GetHandlers(c).Handler.MessageCreate)
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

	registerCommands(dg, c)

	fmt.Println("Bot está ahora en funcionamiento. Presiona Ctrl+C para cerrar.")

	// Espera hasta que se reciba una señal para cerrar el bot
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}

func registerCommands(s *discordgo.Session, c config.Config) {

	guilds := s.State.Guilds
	for _, guild := range guilds {
		_, err := s.ApplicationCommandBulkOverwrite(c.AppId, guild.ID, utils.Commands)
		if err != nil {
			log.Println(err)
		}
	}
}
