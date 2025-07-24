package logger

import (
	"ims/config"
	"os"

	"golang.org/x/exp/slog"
)

var LogLevel = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

func New(cfg *config.Config) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level:     LogLevel[cfg.LogLevel],
		AddSource: true,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	return logger
}
