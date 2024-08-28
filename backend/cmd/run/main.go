package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/GenerateNU/cooked/backend/internal/server"
	"github.com/GenerateNU/cooked/backend/internal/settings"
	"github.com/GenerateNU/cooked/backend/internal/storage/postgres"
)

func main() {
	settings, err := settings.Load("config")
	if err != nil {
		log.Fatal(err)
	}

	db := postgres.New(settings.DB)
	app := server.InitApp(server.Params{Storage: db})

	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", settings.Application.Port)); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	slog.Info("Shutting down server")
	if err := app.Shutdown(); err != nil {
		slog.Error("failed to shutdown server", "error", err)
	}

	slog.Info("Server shutdown")
}
