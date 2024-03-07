package infra

import (
	"BotDiscordGO/internal/application/infra/domain"
	"BotDiscordGO/internal/application/infra/repo"
	"BotDiscordGO/internal/server/infra/config"
)

func GetHandlers(c config.Config) *domain.Handler {
	return &domain.Handler{
		Handler: initHandler(c),
	}
}

func initHandler(c config.Config) *repo.Messages {
	return &repo.Messages{
		Config: &c,
	}
}
