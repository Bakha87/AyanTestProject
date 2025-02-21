package service

import (
	"AyanTestProject/config"
	"AyanTestProject/internal/model"
	"AyanTestProject/internal/repository"
	"context"
)

type Service interface {
	EventNotification(ctx context.Context, event *model.Event) error
	NotificationWorker()
	CloseNotificationWorker()
}

type service struct {
	repository       repository.Repository
	config           *config.Configuration
	SendNotification chan *model.Event
}

func New(repo repository.Repository, cfg *config.Configuration) Service {
	return &service{
		repository:       repo,
		config:           cfg,
		SendNotification: make(chan *model.Event, 1),
	}
}
