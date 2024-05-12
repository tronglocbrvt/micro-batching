package batcher

import (
	"testing"
	"time"

	"github.com/tronglocbrvt/micro-batching/job"
)

// MockBatchProcessor is a mock implementation of BatchProcessor for testing purposes.
type MockBatchProcessor struct{}

func (p *MockBatchProcessor) ProcessBatch(jobs []job.Job) []job.JobResult {
	// Mock implementation
	results := make([]job.JobResult, len(jobs))
	for i, j := range jobs {
		results[i] = job.JobResult{
			JobIDs: []int{j.ID},
			Errors: nil,
		}
	}
	return results
}

// MockLogger is a mock implementation of Logger for testing purposes.
type MockLogger struct {
	InfoLogs  []string
	ErrorLogs []string
}

func (l *MockLogger) Info(message string) {
	l.InfoLogs = append(l.InfoLogs, message)
}

func (l *MockLogger) Error(err error, message string) {
	l.ErrorLogs = append(l.ErrorLogs, message)
}

func TestMicroBatcher_SubmitJob(t *testing.T) {
	logger := &MockLogger{}
	processor := &MockBatchProcessor{}
	batchSize := 2
	batchTimeout := time.Millisecond * 100
	b := NewMicroBatcher(batchSize, batchTimeout, logger, processor)
	go b.Start()

	jobs := []job.Job{
		{ID: 1, Data: "data1"},
		{ID: 2, Data: "data2"},
	}

	for _, j := range jobs {
		b.SubmitJob(j)
	}

	time.Sleep(time.Millisecond * 200) // Wait for processing

	if len(logger.InfoLogs) != 2 {
		t.Errorf("Expected 2 info logs, got %d", len(logger.InfoLogs))
	}
	if len(logger.ErrorLogs) != 0 {
		t.Errorf("Expected 0 error logs, got %d", len(logger.ErrorLogs))
	}
}

func TestMicroBatcher_Shutdown(t *testing.T) {
	logger := &MockLogger{}
	processor := &MockBatchProcessor{}
	batchSize := 2
	batchTimeout := time.Millisecond * 100
	b := NewMicroBatcher(batchSize, batchTimeout, logger, processor)
	go b.Start()

	// Shutdown immediately
	b.Shutdown()

	// Submit a job after shutdown (this will panic)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but got none")
		}
	}()
	b.SubmitJob(job.Job{ID: 1, Data: "data1"})

	// Wait for a short time to allow for processing
	time.Sleep(time.Millisecond * 200)

	// Check that no logs were recorded after shutdown
	if len(logger.InfoLogs) != 0 {
		t.Errorf("Expected 0 info logs, got %d", len(logger.InfoLogs))
	}
	if len(logger.ErrorLogs) != 0 {
		t.Errorf("Expected 0 error logs, got %d", len(logger.ErrorLogs))
	}
}
