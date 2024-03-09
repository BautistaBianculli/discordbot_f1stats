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

// InitApp start application and settings
func InitApp(c config.Config) {

	dg, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		fmt.Println("Error al crear la sesión de DiscordGo,", err)
		return
	}

	dg.AddHandler(GetHandlers(c).Handler.MessageCreate)
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

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

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}

// registerCommands register the commands to the discord bot to be suggested in the different guilds
func registerCommands(s *discordgo.Session, c config.Config) {

	guilds := s.State.Guilds
	for _, guild := range guilds {
		_, err := s.ApplicationCommandBulkOverwrite(c.AppId, guild.ID, utils.Commands)
		if err != nil {
			log.Println(err)
		}
	}
}
