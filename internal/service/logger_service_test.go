package service_test

import (
	"log/slog"
	"testing"

	"example.com/go-auth-globo/internal/service"
)

func TestLogger(t *testing.T) {
	t.Run("it should use global logger", func(t *testing.T) {
		service.InitLoggerService()

		slog.Info("Hello Test")
	})
}
