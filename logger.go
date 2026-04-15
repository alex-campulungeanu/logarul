package logarul

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"gopkg.in/natefinch/lumberjack.v2"
)

func New(cfg Config) (*slog.Logger, *lumberjack.Logger) {
	// FilePath:      "app.log",
	// Level:         slog.LevelDebug,
	// MaxSize:       10,
	// MaxBackups:    3,
	// MaxAge:        7,
	// Compress:      true,
	// AddSource:     true,
	// EnableFile:    true,
	// EnableConsole: true,

	var handlers []slog.Handler
	var rotatingFile *lumberjack.Logger

	if cfg.EnableConsole {
		consoleHandler := tint.NewHandler(os.Stdout, &tint.Options{
			Level:      cfg.Level,
			TimeFormat: time.DateTime,
			AddSource:  cfg.AddSource,
		})
		handlers = append(handlers, consoleHandler)
	}

	if cfg.EnableFile && cfg.FilePath != "" {
		rotatingFile := &lumberjack.Logger{
			Filename:   cfg.FilePath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}

		fileHandler := slog.NewJSONHandler(rotatingFile, &slog.HandlerOptions{AddSource: true})
		handlers = append(handlers, fileHandler)
	}

	// Fallback safety
	if len(handlers) == 0 {
		consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: cfg.Level,
		})
		handlers = append(handlers, consoleHandler)
	}

	tee := &TeeHandler{
		handlers: handlers,
	}

	logger := slog.New(tee)
	return logger, rotatingFile

}
