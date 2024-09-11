package main

import (
	"log/slog"
	"os"
)

func main() {
	level := new(slog.LevelVar)
	logger := (slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource: true,
			Level:     level,
		}))
	slog.SetDefault(slog.New(logger))
	slog.Debug("Lol")
	slog.Info("Lol")
	level.Set(slog.LevelDebug)
	slog.Debug("xd")
	slog.Info("xd")
}
