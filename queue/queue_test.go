package queue

import (
	"github.com/tronglocbrvt/micro-batching/job"
	"testing"
)

// TestQueue tests the Queue functionality.
func TestQueue(t *testing.T) {
	// Initialize Queue
	queueInit := NewQueue()

	// Enqueue jobs
	queueInit.Enqueue(job.Job{ID: 1, Data: "data1"})
	queueInit.Enqueue(job.Job{ID: 2, Data: "data2"})

	// Test Peek
	job, err := queueInit.Peek()
	if err != nil {
		t.Errorf("Error peeking from queue: %v", err)
	}
	if job.ID != 1 || job.Data != "data1" {
		t.Errorf("Unexpected job peeked from queue: %+v", job)
	}

	// Test Dequeue
	job, err = queueInit.Dequeue()
	if err != nil {
		t.Errorf("Error dequeuing from queue: %v", err)
	}
	if job.ID != 1 || job.Data != "data1" {
		t.Errorf("Unexpected job dequeued from queue: %+v", job)
	}

	// Test Size
	size := queueInit.Size()
	if size != 1 {
		t.Errorf("Unexpected queue size: %d", size)
	}

	// Dequeue the remaining job
	_, err = queueInit.Dequeue()
	if err != nil {
		t.Errorf("Error dequeuing from queue: %v", err)
	}

	// Test Size after dequeueing all jobs
	size = queueInit.Size()
	if size != 0 {
		t.Errorf("Unexpected queue size: %d", size)
	}
}
