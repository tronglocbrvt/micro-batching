package logger

import (
	"log"
)

// Logger defines the interface for logging.
type Logger interface {
	Info(message string)
	Error(err error, message string)
}

// DefaultLogger is a simple logger that prints messages to stdout.
type DefaultLogger struct{}

// Info logs an informational message.
func (l *DefaultLogger) Info(message string) {
	log.Println("[INFO]", message)
}

// Error logs an error message along with the error.
func (l *DefaultLogger) Error(err error, message string) {
	log.Printf("[ERROR] %s: %v\n", message, err)
}
