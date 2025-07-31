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

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataBase, err := pgxpool.Connect(
		ctx, fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.DbHost, cfg.DbPort, cfg.DbUser,
			cfg.DbPassword, cfg.DbName, cfg.SSLMode))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dataBase.Close()

	var srv *http.Server

	go func() {
		srv, err = server.Run(ctx, dataBase, cfg)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-ch

	logger.Info("Shutting down server")

	cancel()

	err = server.Close(ctx, dataBase, srv)
	if err != nil {
		logger.Error("Failed to close server", err)
		return
	}
	logger.Info("Server shutdown complete")
}
