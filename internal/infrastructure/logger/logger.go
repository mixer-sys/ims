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
	if cfg.Server.LogLevel != "DEBUG" || cfg.Server.LogLevel != "INFO" || cfg.Server.LogLevel != "WARN" || cfg.Server.LogLevel != "ERROR" {
		cfg.Server.LogLevel = "INFO"
	}

	opts := &slog.HandlerOptions{
		Level: logLevel[cfg.Server.LogLevel],
	}

	if cfg.Server.LogLevel == "DEBUG" {
		opts.AddSource = true
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	return logger
}
