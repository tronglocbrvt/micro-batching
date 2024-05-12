package batcher

import (
	"github.com/tronglocbrvt/micro-batching/job"
	"github.com/tronglocbrvt/micro-batching/logger"
	"github.com/tronglocbrvt/micro-batching/processor"
	"sync"
	"time"
)

// MicroBatcher represents a micro-batching library.
type MicroBatcher struct {
	batchSize    int
	batchTimeout time.Duration
	processor    processor.BatchProcessor
	queue        chan job.Job // Use channel to simulate queue
	shutdown     chan struct{}
	logger       logger.Logger
	cancelMap    map[int]chan struct{}
	wg           sync.WaitGroup
}

// NewMicroBatcher creates a new MicroBatcher instance.
func NewMicroBatcher(batchSize int, batchTimeout time.Duration, logger logger.Logger, processor processor.BatchProcessor) *MicroBatcher {
	return &MicroBatcher{
		batchSize:    batchSize,
		batchTimeout: batchTimeout,
		processor:    processor,
		queue:        make(chan job.Job), // Initialize the channel
		shutdown:     make(chan struct{}),
		logger:       logger,
		cancelMap:    make(map[int]chan struct{}),
	}
}

// Start starts processing jobs in batches.
func (m *MicroBatcher) Start() {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		var batch []job.Job
		var timer *time.Timer

		for {
			select {
			case job := <-m.queue:
				batch = append(batch, job)
				if len(batch) >= m.batchSize {
					m.processBatch(batch)
					batch = nil
					if timer != nil {
						timer.Stop()
						timer = nil
					}
				} else if timer == nil {
					timer = time.AfterFunc(m.batchTimeout, func() {
						m.processBatch(batch)
						batch = nil
					})
				}
			case <-m.shutdown:
				if len(batch) > 0 {
					m.processBatch(batch)
				}
				return
			}
		}
	}()
}

// SubmitJob submits a job to the MicroBatcher.
func (m *MicroBatcher) SubmitJob(job job.Job) {
	m.queue <- job
}

// Shutdown shuts down the MicroBatcher.
func (m *MicroBatcher) Shutdown() {
	close(m.queue)
	close(m.shutdown)
	m.wg.Wait()
}

// processBatch processes a batch of jobs.
func (m *MicroBatcher) processBatch(batch []job.Job) {
	// Process the batch of jobs using the batch processor
	results := m.processor.ProcessBatch(batch)

	// Handle the results, such as logging or error handling
	for _, result := range results {
		// Log job results
		if len(result.Errors) > 0 {
			// Handle errors if present
			for _, err := range result.Errors {
				m.logger.Error(err, "Error processing job")
			}
		} else {
			// Log success if no errors
			m.logger.Info("Job processed successfully")
		}

		// Handle individual job results as needed
	}
}
