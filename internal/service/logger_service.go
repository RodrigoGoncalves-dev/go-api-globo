package service

import (
	"log/slog"
	"os"
	"sync"
)

var (
	logger *slog.Logger
	once   sync.Once
)

func InitLoggerService() {
	once.Do(func() {
		handler := slog.New(slog.NewJSONHandler(os.Stdout, nil))

		slog.SetDefault(handler)
	})
}

func Logger() *slog.Logger {
	if logger == nil {
		InitLoggerService()
	}

	return logger
}
