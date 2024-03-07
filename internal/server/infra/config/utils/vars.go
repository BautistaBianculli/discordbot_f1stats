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
		{
			Name:        "drivertable",
			Description: "Returns the table of the year you sent",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "year",
					Description: "The year for which you want the table",
					Required:    true,
				},
			},
		},
	}
)
