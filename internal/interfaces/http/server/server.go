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
	// Создание пула соединений с базой данных
	dataBase, err := pgxpool.Connect(ctx, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.SSLMode))
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Закрытие пула соединений при завершении контекста
	go func() {
		<-ctx.Done()
		dataBase.Close()
	}()

	r := router.NewRouter(dataBase)

	address := ":" + cfg.ServerPort
	srv := &http.Server{
		Addr:              address,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Запуск HTTP сервера
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server listen error: ",
				slog.String("error", err.Error()),
				slog.String("address", address))
			return
		}
	}()

	// Ожидание завершения контекста
	<-ctx.Done()

	// Завершение работы сервера
	if err := srv.Shutdown(context.Background()); err != nil {
		slog.Error("server shutdown error: ",
			slog.String("error", err.Error()),
			slog.String("address", address))
		return fmt.Errorf("server shutdown error: %w", err)
	}

	return nil
}
