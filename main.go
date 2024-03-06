package main

import (
	"BotDiscordGO/internal/server/infra/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const (
	prefix       = "!"
	endOfMessage = "FIUUUUUUUUUUUUUUUUUUUUUUM"
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
	dg.AddHandler(messageCreate)

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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignora los mensajes del propio bot
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, prefix) {
		command := strings.ToLower(strings.TrimSpace(m.Content[len(prefix):])) // Elimina el prefijo y convierte a minúsculas
		switch command {
		case "who":
			_, _ = s.ChannelMessageSend(m.ChannelID, "¡Hi "+m.Author.Mention()+"! I'm Kevin Schumacher, i'm still in development by my creator, be patient so I can bring you the best Formula 1 statistics. "+endOfMessage)
		}
	}
}
