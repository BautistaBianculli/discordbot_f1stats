package repo

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

const (
	prefix        = "!"
	endOfMessage  = "FIUUUUUUUUUUUUUUUUUUUUUUM"
	readerMessage = "This is my image. "
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignora los mensajes del propio bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	log.Printf("The user %s have the id %s", m.Author.Mention(), m.Author.ID)
	if strings.HasPrefix(m.Content, prefix) {
		command := strings.ToLower(strings.TrimSpace(m.Content[len(prefix):])) // Elimina el prefijo y convierte a minúsculas
		switch command {
		case "who":
			whoMessage(s, m)
		case "avatar":
			avatarMessage(s, m)
		}

	}
}

func whoMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "¡Hi "+m.Author.Mention()+"! I'm Kevin Schumacher, i'm still in development by my creator, be patient so I can bring you the best Formula 1 statistics. "+endOfMessage)
}

func avatarMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelFileSend(m.ChannelID, "../../../assets/image.png", strings.NewReader(readerMessage+endOfMessage))
}
