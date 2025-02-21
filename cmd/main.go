package main

import (
	"AyanTestProject/config"
	"AyanTestProject/internal/handler"
	"AyanTestProject/internal/repository"
	"AyanTestProject/internal/server"
	"AyanTestProject/internal/service"
	"AyanTestProject/pkg/database/postgres"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.New("config.yml")
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	pg, err := postgres.New(cfg.PG.URL,
		postgres.MaxPoolSize(cfg.PG.PoolMax),
		postgres.ConnAttempts(cfg.PG.ConnAttempts))
	if err != nil {
		slog.Error("Can not create connect with pg. Error: " + err.Error())
		return
	}
	defer pg.Close()

	// Инициализация репозитория, сервиса, хендлера и сервера
	repo := repository.New(pg)

	srv := service.New(repo, cfg)

	go srv.NotificationWorker()

	defer srv.CloseNotificationWorker()

	h := handler.New(srv)

	src := server.New(h)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go src.Run()

	<-quit
	slog.Info("Shutting down server...")
}
