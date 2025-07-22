package main

import (
	"database/sql"
	"fmt"
	config "ims/config"
	"ims/internal/infrastructure/adapters/logger"
	router "ims/internal/infrastructure/adapters/router"

	"net/http"
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

	logger := logger.GetLogger(cfg)

	dataBase, err := sql.Open(cfg.GooseDriver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.SSLMode))
	if err != nil {
		logger.Error("Failed to connect to the database", err)
		os.Exit(1)
	}
	defer dataBase.Close()

	r := router.NewRouter(dataBase)
	address := ":" + cfg.ServerPort
	logger.Info("Starting server on :", address)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	if err := http.ListenAndServe(address, r); err != nil {
		logger.Error("ListenAndServe error", err)
	}
	<-ch
	logger.Info("Shutting down server")
	if err := dataBase.Close(); err != nil {
		logger.Error("Failed to close database connection", err)
	}
	logger.Info("Server gracefully stopped")

}
