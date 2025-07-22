package server

import (
	"context"
	"database/sql"
	"fmt"
	"ims/config"
	"log/slog"
	"net/http"
	"time"

	router "ims/internal/infrastructure/adapters/router"
)

func Run(ctx context.Context, cfg *config.Config) error {

	dataBase, err := sql.Open(cfg.GooseDriver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.SSLMode))
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	go func() {
		<-ctx.Done()
		if err := dataBase.Close(); err != nil {
			slog.Error("failed to close database connection: ",
				slog.String("error", err.Error()))
			return
		}
	}()

	r := router.NewRouter(dataBase)

	address := ":" + cfg.ServerPort
	srv := &http.Server{
		Addr:              address,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Log the error or handle it as needed
			slog.Error("server listen error: ",
				slog.String("error", err.Error()),
				slog.String("address", address))
			return
		}
	}()

	<-ctx.Done()

	if err := srv.Shutdown(context.Background()); err != nil {
		slog.Error("server shutdown error: ",
			slog.String("error", err.Error()),
			slog.String("address", address))
		return fmt.Errorf("server shutdown error: %w", err)
	}

	return nil
}
