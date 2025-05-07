package main

import (
	"log/slog"
)

type Application struct {
	Logger *slog.Logger
}

func (app *Application) NewApplication(logger *slog.Logger) *Application {
	return &Application{
		Logger: logger,
	}
}
