package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
)

// Logger is the main logging structure that aggregates one or more log writers.
type Logger struct {
	writer io.Writer
}

// LoggerConfig holds configuration options for creating a Logger.
type LoggerConfig struct {
	// ConsoleWriter indicates whether to output logs to the console.
	ConsoleWriter bool

	// FileWriter is an optional file writer destination (implements FileWriter interface).
	FileWriter FileWriter

	// Level sets the global logging level (e.g., zerolog.InfoLevel, zerolog.DebugLevel).
	Level zerolog.Level
}

// New creates and returns a new Logger instance based on the provided configuration.
// It returns an error if no log writer is configured.
func New(cfg *LoggerConfig) (*Logger, error) {
	// Set global time format for logs.
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var writers []io.Writer

	// Append console writer if enabled.
	if cfg.ConsoleWriter {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Append file writer if provided.
	if cfg.FileWriter != nil {
		writers = append(writers, cfg.FileWriter)
	}

	// Ensure at least one writer is available.
	if len(writers) == 0 {
		return nil, fmt.Errorf("at least one writer (console or file) must be configured")
	}

	// Combine writers into a multi-writer.
	multiWriter := io.MultiWriter(writers...)

	// Set the global log level.
	zerolog.SetGlobalLevel(cfg.Level)

	return &Logger{writer: multiWriter}, nil
}

// Sub creates a SubLogger with an associated topic. The topic is automatically
// included in every log entry via the zerolog context.
func (l *Logger) Sub(topic string) *SubLogger {
	// Create a new zerolog logger with the topic field and a timestamp.
	zlogger := zerolog.New(l.writer).With().Str("topic", topic).Timestamp().Logger()
	return NewSubLogger(zlogger)
}
