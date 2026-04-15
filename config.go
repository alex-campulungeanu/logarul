package logarul

import "log/slog"

type Config struct {
	FilePath      string
	Level         slog.Level
	MaxSize       int
	MaxBackups    int
	MaxAge        int
	Compress      bool
	AddSource     bool
	EnableFile    bool
	EnableConsole bool
}

func NewMinimalConfig() Config {
	return Config{
		MaxSize:       5,
		MaxBackups:    1,
		MaxAge:        1,
		Compress:      true,
		AddSource:     true,
		EnableFile:    true,
		EnableConsole: true,
	}
}
