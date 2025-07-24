package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
)

// Logger wraps zerolog.Logger and handles file closing
type Logger struct {
	zerolog.Logger
	logFile *os.File
}

// New creates a new Logger instance with separate levels for console and file
func New(consoleLevel, fileLevel string) (*Logger, error) {
	parsedConsoleLevel, err := zerolog.ParseLevel(consoleLevel)
	if err != nil {
		return nil, fmt.Errorf("invalid console log level: %w", err)
	}
	parsedFileLevel, err := zerolog.ParseLevel(fileLevel)
	if err != nil {
		return nil, fmt.Errorf("invalid file log level: %w", err)
	}

	// Set global minimum level (lowest of both)
	minLevel := parsedConsoleLevel
	if parsedFileLevel < minLevel {
		minLevel = parsedFileLevel
	}
	zerolog.SetGlobalLevel(minLevel)

	// Console writer (human-readable with color)
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "15:04:05",
	}
	consoleWriter.FormatLevel = func(i interface{}) string {
		return fmt.Sprintf("\x1b[36m%-6s\x1b[0m", i) // Cyan level
	}
	consoleWriter.FormatTimestamp = func(i interface{}) string {
		return fmt.Sprintf("\x1b[90m%s\x1b[0m", i) // Gray timestamp
	}

	// Ensure log directory exists
	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	// File writer (JSON)
	logFilePath := filepath.Join(logDir, "app.log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	// Combine loggers using MultiLevelWriter
	multi := zerolog.MultiLevelWriter(consoleWriter, logFile)
	baseLogger := zerolog.New(multi).With().Timestamp().Logger().Level(minLevel)

	return &Logger{
		Logger:  baseLogger,
		logFile: logFile,
	}, nil
}

// Close safely closes the log file
func (l *Logger) Close() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}