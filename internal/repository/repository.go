package repository

import (
	"AyanTestProject/internal/model"
	"AyanTestProject/pkg/database/postgres"
	"context"
)

const (
	timeout = 3
)

type Repository interface {
	PutEventNotification(ctx context.Context, req *model.Event) error
}

type repository struct {
	pg *postgres.Postgres
}

func New(pg *postgres.Postgres) Repository {
	return &repository{pg: pg}
}
