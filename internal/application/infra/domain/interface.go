package domain

import "github.com/bwmarrin/discordgo"

type MessageHandler interface {
	MessageCreate(s *discordgo.Session, i *discordgo.InteractionCreate)
}
