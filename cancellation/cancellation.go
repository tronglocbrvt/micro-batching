package cancellation

import (
	"sync"
)

// JobCancellation represents the cancellation mechanism for jobs.
type JobCancellation struct {
	mu         sync.Mutex
	cancelFunc map[int]func()
}

// NewJobCancellation creates a new JobCancellation instance.
func NewJobCancellation() *JobCancellation {
	return &JobCancellation{
		cancelFunc: make(map[int]func()),
	}
}

// RegisterCancellation registers a cancellation function for the given job ID.
func (c *JobCancellation) RegisterCancellation(jobID int, cancelFunc func()) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cancelFunc[jobID] = cancelFunc
}

// Cancel cancels the job with the given ID.
func (c *JobCancellation) Cancel(jobID int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cancelFunc, ok := c.cancelFunc[jobID]; ok {
		cancelFunc()
		delete(c.cancelFunc, jobID)
	}
}
