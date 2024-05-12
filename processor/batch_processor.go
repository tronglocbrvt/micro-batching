package processor

import (
	"github.com/tronglocbrvt/micro-batching/job"
)

// BatchProcessor defines the interface for processing batches of jobs.
type BatchProcessor interface {
	ProcessBatch(jobs []job.Job) []job.JobResult
}

// DefaultBatchProcessor is a simple batch processor that processes jobs sequentially.
type DefaultBatchProcessor struct{}

// ProcessBatch processes a batch of jobs sequentially.
func (p *DefaultBatchProcessor) ProcessBatch(jobs []job.Job) []job.JobResult {
	results := make([]job.JobResult, len(jobs))
	for i, job := range jobs {
		// Process the job
		result := p.processJob(job)
		results[i] = result
	}
	return results
}

// processJob simulates processing of a single job and returns a job result.
func (p *DefaultBatchProcessor) processJob(j job.Job) job.JobResult {
	// Simulate processing by just returning a successful result
	return job.JobResult{
		JobIDs: []int{j.ID},
		Errors: nil,
		Data:   []interface{}{j.Data},
	}
}
