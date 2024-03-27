// package logger would be implemented from a go-kit repository
package logger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/cjsampson/weather-service/app/conf"
)

// New create simple logger for application
func New(c conf.LogConfig) *slog.Logger {
	level := NewStringLeveler(c.LogLevel)

	opts := &slog.HandlerOptions{
		AddSource: false,
		Level:     level,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	return slog.New(handler)
}

func NewStringLeveler(levelStr string) *StringLeveler {
	level := mapEnvLevelToSlogLevel(levelStr)

	return &StringLeveler{level: level}
}

// mapEnvLevelToSlogLevel mapping function for string environment variable
func mapEnvLevelToSlogLevel(levelStr string) slog.Level {
	switch strings.ToLower(levelStr) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

type StringLeveler struct {
	level slog.Level
}

func (l StringLeveler) Level() slog.Level {
	return l.level
}
