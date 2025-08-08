package server

import (
	"context"
	"errors"
	"fmt"
	"ims/config"
	"ims/internal/domain/repository"
	router "ims/internal/infrastructure/adapters/router"
	"ims/internal/infrastructure/interfaces/http/middleware"
	"log/slog"
	"net/http"
	"time"
)

func Run(ctx context.Context,
	dataBase repository.Database, cfg *config.Config) (
	*http.Server, error) {
	r := router.NewRouter(dataBase)

	address := ":" + cfg.Server.Port

	srv := &http.Server{
		Addr:              address,
		Handler:           middleware.LoggerMiddleware(r),
		ReadHeaderTimeout: time.Duration(cfg.ReadHeaderTimeoutSecond) * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server listen error: ",
				slog.String("error", err.Error()),
				slog.String("address", address))

			return
		}
	}()

	return srv, nil
}

func Close(ctx context.Context,
	srv *http.Server) error {

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error: ",
			slog.String("error", err.Error()),
		)

		return fmt.Errorf("server shutdown error: %w", err)
	}

	return nil
}
