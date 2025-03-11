// `log` package for free form output
// `slog` package for structured output

package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	// reasonable logging to `os.Stderr`
	// `Fatal*` and `Panic*` exit program after logging
	log.Println("standard logger")

	// loggers can be configured with flags to set their output format
	// standard logger has `log.Ldate` and `log.Ltime` flags set
	// example: change to emit time eith microsecond accuracy
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with microseconds")

	// emit filename and line from which the `log` function is called
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// create a custom logger and pass it around
	// we can set a prefix to distinguish its output from other loggers
	myLog := log.New(os.Stdout, "myLog:", log.LstdFlags)
	myLog.Println("from my log")

	// set prefix on existing loggers
	myLog.SetPrefix("ohmy:")
	myLog.Println("from my log")

	// loggers can have custom output targets; any `io.Writer` works
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// write output into buf
	buflog.Println("hello")
	fmt.Println("from buflog:", buf.String())

	// structured log output
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// `slog` output can contain any arbitrary number of key=value pairs
	myslog.Info("hello again", "a", "1", "b", 2)

	// text logs
	fmt.Println()
	textHandler := slog.NewTextHandler(os.Stderr, nil)
	myslogText := slog.New(textHandler)
	myslogText.Info("hi there")
	myslogText.Info("hi there again", "a", "1", "b", 2)
}
