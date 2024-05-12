# Micro-Batching Library
The Micro-Batching Library is a Go package that provides utilities for micro-batching of jobs for efficient processing. It allows you to group jobs into batches and process them in parallel, improving throughput and resource utilization.

## Features
### Batch Processing:
_Group jobs into batches for efficient processing._
- **Concurrency**: Utilize concurrency to process batches in parallel.
- **Customizable**: Customize batch size, batch timeout, and batch processing logic.
- **Logging**: Integration with logging framework for tracking job processing and errors.
- **Cancellation**: Mechanism to cancel individual jobs or entire batches.
## Installation
To use the Micro-Batching Library in your Go project, you can simply install it using go get:
```bash
go get github.com/tronglocbrvt/micro-batching
```

## Usage
Here's a basic example of how to use the Micro-Batching Library:
```bash
package main

import (
    "github.com/tronglocbrvt/micro-batching/batcher"
    "github.com/tronglocbrvt/micro-batching/logger"
    "github.com/tronglocbrvt/micro-batching/processor"
    "time"
)

func main() {
    // Initialize logger
    logger := &logger.DefaultLogger{}

    // Initialize batch processor
    processor := &processor.DefaultBatchProcessor{}

    // Initialize MicroBatcher
    batchSize := 10
    batchTimeout := 1 * time.Second
    batcher := batcher.NewMicroBatcher(batchSize, batchTimeout, logger, processor)
    batcher.Start()

    // Submit jobs
    for i := 0; i < 100; i++ {
        batcher.SubmitJob(job.Job{ID: i, Data: "data"})
    }

    // Shutdown when finish
    batcher.Shutdown()
}
```