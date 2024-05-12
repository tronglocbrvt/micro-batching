package queue

import (
	"errors"
	"github.com/tronglocbrvt/micro-batching/job"
	"sync"
)

// Queue represents a sophisticated queue for storing jobs.
type Queue struct {
	items []job.Job
	mu    sync.Mutex
}

// NewQueue creates a new Queue instance.
func NewQueue() *Queue {
	return &Queue{
		items: make([]job.Job, 0),
	}
}

// Enqueue adds a job to the end of the queue.
func (q *Queue) Enqueue(job job.Job) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, job)
}

// Dequeue removes and returns the first job from the queue.
func (q *Queue) Dequeue() (job.Job, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return job.Job{}, errors.New("queue is empty")
	}
	job := q.items[0]
	q.items = q.items[1:]
	return job, nil
}

// Peek returns the first job in the queue without removing it.
func (q *Queue) Peek() (job.Job, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return job.Job{}, errors.New("queue is empty")
	}
	return q.items[0], nil
}

// Size returns the number of jobs in the queue.
func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}
