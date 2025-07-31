package main

import (
	"context"
	"fmt"
	config "ims/config"
	"ims/internal/infrastructure/logger"
	"ims/internal/interfaces/http/server"
	"log"
	"net/http"

	"os"
	"os/signal"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	logger := logger.New(cfg)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataBase, err := pgxpool.Connect(
		ctx, fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.DBHost, cfg.DBPort, cfg.DBUser,
			cfg.DBPassword, cfg.DBName, cfg.SSLMode))
	if err != nil {
		logger.Error("Failed to connect to database", err)

		return
	}

	defer dataBase.Close()

	var srv *http.Server

	go func() {
		srv, err = server.Run(ctx, dataBase, cfg)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-shutdown

	logger.Info("Shutting down server")

	cancel()

	err = server.Close(ctx, dataBase, srv)
	if err != nil {
		logger.Error("Failed to close server", err)

		return
	}

	logger.Info("Server shutdown complete")
}
