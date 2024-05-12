package job

// Job represents a task to be processed.
type Job struct {
	ID   int
	Data interface{}
}

// JobResult represents the result of processing a batch of jobs.
type JobResult struct {
	JobIDs []int
	Errors []error
	Data   []interface{}
}
