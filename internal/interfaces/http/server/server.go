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

func Run(ctx context.Context, dataBase *pgxpool.Pool, cfg *config.Config) (*http.Server, error) {

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

	return srv, nil
}

func Close(ctx context.Context, dataBase *pgxpool.Pool, srv *http.Server) error {

	dataBase.Close()

	if err := srv.Shutdown(context.Background()); err != nil {
		slog.Error("server shutdown error: ",
			slog.String("error", err.Error()),
		)
		return fmt.Errorf("server shutdown error: %w", err)
	}
	return nil
}
