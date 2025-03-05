package logger

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

// FileWriter defines an interface for file-based log writing.
// It is simply an alias for io.Writer and allows future extension.
type FileWriter interface {
	io.Writer
}

// NewFileWriter creates a new file writer using the lumberjack package for log rotation.
// Parameters:
//   - filename: Path to the log file.
//   - maxSize: Maximum size in MB before the log rotates.
//   - maxBackups: Maximum number of rotated backup files to retain.
//   - compress: Whether to compress rotated log files.
func NewFileWriter(filename string, maxSize int, maxBackups int, compress bool) FileWriter {
	return &lumberjack.Logger{
		Filename:   filename,    // Path to the log file.
		MaxSize:    maxSize,     // Max size in megabytes before rotation.
		MaxBackups: maxBackups,  // Maximum number of backup files.
		Compress:   compress,    // Compress rotated files if true.
	}
}
