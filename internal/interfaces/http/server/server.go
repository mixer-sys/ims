package server

import (
	"context"
	"fmt"
	"ims/config"
	"log/slog"
	"net/http"
	"time"

	router "ims/internal/infrastructure/adapters/router"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Run(ctx context.Context, cfg *config.Config) error {

	dataBase, err := pgxpool.Connect(
		ctx, fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.DbHost, cfg.DbPort, cfg.DbUser,
			cfg.DbPassword, cfg.DbName, cfg.SSLMode))
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	defer dataBase.Close()

	r := router.NewRouter(dataBase)

	address := ":" + cfg.ServerPort

	srv := &http.Server{
		Addr:              address,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server listen error: ",
				slog.String("error", err.Error()),
				slog.String("address", address))
			return
		}
	}()

	Close(ctx, dataBase, srv)

	return nil
}

func Close(ctx context.Context, dataBase *pgxpool.Pool, srv *http.Server) error {
	go func() {
		<-ctx.Done()
		dataBase.Close()
	}()

	<-ctx.Done()

	if err := srv.Shutdown(context.Background()); err != nil {
		slog.Error("server shutdown error: ",
			slog.String("error", err.Error()),
		)
		return fmt.Errorf("server shutdown error: %w", err)
	}
	return nil
}
