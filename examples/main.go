package main

import (
	"errors"
	"log"

	"github.com/rs/zerolog"
	"github.com/tahadostifam/logger"
)

func main() {
	// Configure the logger with both console and file writers,
	// and set the log level to Info.
	cfg := &logger.LoggerConfig{
		ConsoleWriter: true,
		FileWriter:    logger.NewFileWriter("./logs/logs", 2, 1, true),
		Level:         zerolog.InfoLevel,
	}

	// Initialize the main logger using the provided configuration.
	mainLogger, err := logger.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	// Create sub-loggers with specific topics.
	fooLogger := mainLogger.Sub("Foo")
	barLogger := mainLogger.Sub("Bar")

	// Log an error message with additional fields.
	fooLogger.Error("unable to create the foo in database", errors.New("error detail"), map[string]string{
		"name": "Sample key for field of error method",
	})

	// Log a warning message with additional fields.
	fooLogger.Warn("unable to create the foo in database", map[string]string{
		"name": "Sample key for field of warn method",
	})

	// Log an informational message with additional fields.
	fooLogger.Info("hello world hello world hello world", map[string]string{
		"name": "Sample key for field of warn method",
	})

	// Log a debug message with additional fields.
	barLogger.Debug("debug mebug debug mebug debug mebug", map[string]string{
		"name": "Sample key for field of warn method",
	})

	// Log a trace message with additional fields.
	barLogger.Trace("trace mrace trace mrace trace mrace", map[string]string{
		"name": "Sample key for field of warn method",
	})
}
