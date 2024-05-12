package logger

import (
	"errors"
	"log"
	"strings"
	"testing"
)

// TestDefaultLogger_Info tests the Info method of DefaultLogger.
func TestDefaultLogger_Info(t *testing.T) {
	// Create a DefaultLogger instance
	logger := &DefaultLogger{}

	// Capture stdout
	logOutput := captureLogOutput(func() {
		// Call the Info method with a test message
		logger.Info("Test info message")
	})

	// Check if the expected log message is printed
	expectedOutput := "[INFO] Test info message\n"
	if !strings.Contains(logOutput, expectedOutput) {
		t.Errorf("Expected log output to be %q, but got %q", expectedOutput, logOutput)
	}
}

// TestDefaultLogger_Error tests the Error method of DefaultLogger.
func TestDefaultLogger_Error(t *testing.T) {
	// Create a DefaultLogger instance
	logger := &DefaultLogger{}

	// Create a test error
	testErr := errors.New("test error")

	// Capture stdout
	logOutput := captureLogOutput(func() {
		// Call the Error method with a test error and message
		logger.Error(testErr, "Test error message")
	})

	// Check if the expected log message is printed
	expectedOutput := "[ERROR] Test error message: test error\n"
	if !strings.Contains(logOutput, expectedOutput) {
		t.Errorf("Expected log output to be %q, but got %q", expectedOutput, logOutput)
	}
}

// captureLogOutput is a helper function to capture stdout output.
func captureLogOutput(f func()) string {
	var logOutput string
	log.SetOutput(&capture{&logOutput})
	defer func() {
		log.SetOutput(log.Writer())
	}()
	f()
	return logOutput
}

// capture is a custom writer to capture log output.
type capture struct {
	output *string
}

func (c *capture) Write(p []byte) (n int, err error) {
	*c.output += string(p)
	return len(p), nil
}
