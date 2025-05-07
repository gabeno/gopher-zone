package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	app := NewApplication(logger)

	logger.Info("starting server", slog.Any("addr", *addr))

	err := http.ListenAndServe(*addr, app.routes())
	log.Fatal(err)
	logger.Error(err.Error())
	os.Exit(1)
}
