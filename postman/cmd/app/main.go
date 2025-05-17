package main

import (
	"os"
	"os/signal"
	"postman/internal/app"
	"postman/pkg/lib/logger"
	"syscall"
)

func main() {
	log := logger.SetupLogger("local")

	application := app.New(log)

	go func() {
		application.Run()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Info("application stopped")
}
