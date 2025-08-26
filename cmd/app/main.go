package main

import (
	"context"
	"ims/config"
	"ims/internal/domain/repository"
	"ims/internal/infrastructure/interfaces/http/server"
	"ims/internal/infrastructure/logger"
	"log"
	"net/http"

	"os"
	"os/signal"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	logger := logger.New(cfg)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataBase, err := repository.NewDatabase(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	defer dataBase.Close()

	var srv *http.Server

	go func() {
		srv, err = server.Run(ctx, *dataBase, cfg)
		if err != nil {
			log.Fatalf("Failed to start server: %s", err.Error())
		}
	}()

	<-shutdown

	logger.Info("Shutting down server")

	cancel()

	err = server.Close(ctx, srv)
	if err != nil {
		logger.Error("Failed to close server", err)

		return
	}

	logger.Info("Server shutdown complete")
}
