package logger

import (
	"log/slog"
	"os"
	"time"
)

func New(level string) *slog.Logger {
	if level == "prod" {
		return slog.New(createProdLog())
	}

	return slog.New(createDevLog())
}

func createDevLog() *slog.TextHandler {
	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.Int64Value(time.Now().Unix())
			}

			return a
		},
	})
}

func createProdLog() *slog.JSONHandler {
	return slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: false,
	})
}
