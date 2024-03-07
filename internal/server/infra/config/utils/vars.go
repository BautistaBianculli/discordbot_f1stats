package utils

import "github.com/bwmarrin/discordgo"

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "avatar",
			Description: "Returns my avatar image",
		},
		{
			Name:        "who",
			Description: "Returns info of my",
		},
		{
			Name:        "help",
			Description: "Send my commands information",
		},
	}
)
