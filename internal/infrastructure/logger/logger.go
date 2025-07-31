package logger

import (
	"ims/config"
	"os"

	"golang.org/x/exp/slog"
)

func New(cfg *config.Config) *slog.Logger {
	var logLevel = map[string]slog.Level{
		"DEBUG": slog.LevelDebug,
		"INFO":  slog.LevelInfo,
		"WARN":  slog.LevelWarn,
		"ERROR": slog.LevelError,
	}

	opts := &slog.HandlerOptions{
		Level:     logLevel[cfg.LogLevel],
		AddSource: true,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	return logger
}
