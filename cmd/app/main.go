package main

import (
	"database/sql"
	"fmt"
	"ims/internal/config"
	"ims/internal/logger"
	"ims/internal/routers"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load configuration: %v", err)
		os.Exit(1)
	}

	logger := logger.GetLogger(config)

	dataBase, err := sql.Open(config.GooseDriver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName, config.SSLMode))
	if err != nil {
		logger.Error("Failed to connect to the database", err)
		os.Exit(1)
	}
	defer dataBase.Close()

	r := routers.NewRouter(dataBase)
	address := ":" + config.ServerPort
	logger.Info("Starting server on %s", address)
	if err := http.ListenAndServe(address, r); err != nil {
		logger.Error("ListenAndServe error", err)
	}
}
