package repo

import (
	"BotDiscordGO/internal/application/infra/domain"
	"BotDiscordGO/internal/server/infra/config"
	"BotDiscordGO/internal/server/infra/config/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

const (
	endOfMessage    = "FIUUUUUUUUUUUUUUUUUUUUUUM"
	readerMessage   = "This is my image %s. "
	whoMessageConst = "¡Hi %s! I'm Kevin Schumacher, i'm still in development by my creator, be patient so I can bring you the best Formula 1 statistics. "
)

type Messages struct {
	Config *config.Config
}

type MessagePriv interface {
	helpMessage(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func (m *Messages) MessageCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {

	if i.Interaction.ID == s.State.User.ID {
		return
	}
	data := i.ApplicationCommandData()
	log.Printf("The user %s have the id %s", i.Interaction.Member.Nick, i.Interaction.Member.User.ID)
	switch data.Name {
	case "who":
		whoMessage(s, i)
	case "avatar":
		avatarMessage(s, i)
	case "help":
		m.helpMessage(s, i)
	}
}

func (m *Messages) helpMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	helpMessages := make([]domain.HelpMessage, 0)
	for _, command := range utils.Commands {
		helpMessage := domain.HelpMessage{
			Name:    command.Name,
			Content: command.Description,
		}
		helpMessages = append(helpMessages, helpMessage)
	}
	data, _ := s.GuildEmojis(m.Config.GuildMainId)
	helpText := fmt.Sprintf("%s \t**COMMANDS**\n", data[0].MessageFormat())
	for _, helpMessage := range helpMessages {
		helpText += fmt.Sprintf("**/%s** %s\n", helpMessage.Name, helpMessage.Content)
	}

	helpText += "\n" + endOfMessage

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: helpText,
		},
	}
	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		log.Printf("Error: %v", err)
	}
}

func whoMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(whoMessageConst+endOfMessage, i.Interaction.Member.Mention())},
	}

	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		log.Printf("Error: %v", err)
	}
}

func avatarMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	author := discordgo.MessageEmbedAuthor{
		Name:    fmt.Sprintf(readerMessage, i.Interaction.Member.Nick),
		IconURL: discordgo.EndpointUserAvatar(s.State.User.ID, s.State.User.Avatar),
	}

	embed := discordgo.MessageEmbed{
		Title: author.Name,
		Image: &discordgo.MessageEmbedImage{
			URL: author.IconURL,
		},
	}

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&embed},
		},
	}
	err := s.InteractionRespond(i.Interaction, response)

	if err != nil {
		log.Printf("Error: %v", err)
	}
}
