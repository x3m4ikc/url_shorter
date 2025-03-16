package main

import (
	"fmt"
	"log/slog"
	"os"
	"url_shorter/internal/config"
	"url_shorter/internal/lib/logger/sl"
	"url_shorter/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev = "dev"
	endProd = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("[ STARTING URL-SHORTENER ]", slog.String("env", cfg.Env))
	log.Debug("DEBUG LOGS")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	// TODO: init router

	// TODO: init server

	fmt.Println(cfg)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case endProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	
	return log
}