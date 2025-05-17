package main

import (
	"api/internal/app"
	"api/internal/storage/psql"
	"api/pkg/config"
	"api/pkg/lib/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	connStr := "postgres://postgres:123@psql_api:5432/api_db?sslmode=disable"
	storage := psql.New(log, connStr)

	application := app.New(log, cfg.Port, storage)
	go func() {
		application.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	storage.Close()
	log.Info("database stopped")
	log.Info("application stopped")
}
