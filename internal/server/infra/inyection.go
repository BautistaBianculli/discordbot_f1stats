package infra

import (
	"BotDiscordGO/internal/application/infra/domain"
	msgRepo "BotDiscordGO/internal/application/infra/repo"
	f1Repo "BotDiscordGO/internal/f1api/infra/repo"
	"BotDiscordGO/internal/server/infra/config"
)

func GetHandlers(c config.Config) *domain.Handler {
	return &domain.Handler{
		Handler: initHandler(c),
	}
}

func initHandler(c config.Config) *msgRepo.Messages {
	return &msgRepo.Messages{
		Config:      &c,
		FRepository: initF1Repository(c),
	}
}

func initF1Repository(c config.Config) *f1Repo.F1Repository {
	return &f1Repo.F1Repository{
		Config: &c,
		Client: config.NewHttpclient(),
	}
}
