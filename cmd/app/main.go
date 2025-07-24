package main

import (
	"context"
	"fmt"
	config "ims/config"
	"ims/internal/infrastructure/adapters/logger"
	"ims/internal/interfaces/http/server"

	"os"
	"os/signal"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load configuration: %v", err)
		os.Exit(1)
	}

	logger := logger.New(cfg)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				logger.Info("Context cancelled, shutting down server")
				return
			default:
				err := server.Run(ctx, cfg)
				if err != nil {
					logger.Error("Failed to run server", err)
					return
				}
			}
		}
	}()

	<-ch

	logger.Info("Shutting down server")

	cancel()

	logger.Info("Server shutdown complete")
}
