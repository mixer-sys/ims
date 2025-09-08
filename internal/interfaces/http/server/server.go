package server

import (
	"context"
	"errors"
	"fmt"
	"ims/config"
	router "ims/internal/adapters/router"
	"ims/internal/domain/repository"
	"ims/internal/interfaces/http/middleware"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	address    string
}

func New(dataBase *repository.Database, cfg *config.Config) *Server {
	s := &Server{}
	r := router.NewRouter(*dataBase)

	s.address = ":" + cfg.Server.Port
	s.httpServer = &http.Server{
		Addr:              s.address,
		Handler:           middleware.LoggerMiddleware(r),
		ReadHeaderTimeout: time.Duration(cfg.ReadHeaderTimeoutSecond) * time.Second,
	}
	return s
}

func (s *Server) Run(ctx context.Context) error {

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server listen error: ",
				slog.String("error", err.Error()),
				slog.String("address", s.address))

			return
		}
	}()

	return nil
}

func (s *Server) Close(ctx context.Context) error {

	if err := s.httpServer.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error: ",
			slog.String("error", err.Error()),
		)

		return fmt.Errorf("server shutdown error: %w", err)
	}

	return nil
}
