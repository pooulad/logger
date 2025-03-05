package logger

import (
	"github.com/rs/zerolog"
)

// Fields represents a collection of key-value pairs for structured logging.
type Fields map[string]string

// SubLogger provides logging methods with additional context, such as a topic.
type SubLogger struct {
	logger zerolog.Logger
}

// NewSubLogger creates a new SubLogger using the provided zerolog.Logger instance.
// The logger should already have context (e.g., topic, timestamp) set.
func NewSubLogger(zlogger zerolog.Logger) *SubLogger {
	return &SubLogger{logger: zlogger}
}

// addFields adds additional key-value fields to the zerolog event.
func addFields(fields Fields, event *zerolog.Event) {
	for key, value := range fields {
		event.Str(key, value)
	}
}

// Panic logs a message at the Panic level along with an error and additional fields.
func (sl *SubLogger) Panic(reason string, err error, fields Fields) {
	event := sl.logger.Panic().Err(err)
	addFields(fields, event)
	event.Msg(reason)
}

// Error logs a message at the Error level along with an error and additional fields.
func (sl *SubLogger) Error(reason string, err error, fields Fields) {
	event := sl.logger.Error().Err(err)
	addFields(fields, event)
	event.Msg(reason)
}

// Fatal logs a message at the Fatal level along with an error and additional fields.
func (sl *SubLogger) Fatal(reason string, err error, fields Fields) {
	event := sl.logger.Fatal().Err(err)
	addFields(fields, event)
	event.Msg(reason)
}

// Warn logs a warning message with additional fields.
func (sl *SubLogger) Warn(msg string, fields Fields) {
	event := sl.logger.Warn()
	addFields(fields, event)
	event.Msg(msg)
}

// Info logs an informational message with additional fields.
func (sl *SubLogger) Info(msg string, fields Fields) {
	event := sl.logger.Info()
	addFields(fields, event)
	event.Msg(msg)
}

// Debug logs a debug message with additional fields.
func (sl *SubLogger) Debug(msg string, fields Fields) {
	event := sl.logger.Debug()
	addFields(fields, event)
	event.Msg(msg)
}

// Trace logs a trace message with additional fields.
func (sl *SubLogger) Trace(msg string, fields Fields) {
	event := sl.logger.Trace()
	addFields(fields, event)
	event.Msg(msg)
}
