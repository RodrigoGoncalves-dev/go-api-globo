package service

import (
	"log/slog"
	"os"
)

func Logger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return logger
}
