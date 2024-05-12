package processor

import (
	"github.com/tronglocbrvt/micro-batching/job"
	"testing"
)

func TestCustomBatchProcessor_ProcessBatch(t *testing.T) {
	// Create a custom batch processor instance
	processor := &DefaultBatchProcessor{}

	// Create some test jobs
	testJobs := []job.Job{
		{ID: 1, Data: "data1"},
		{ID: 2, Data: "data2"},
	}

	// Call the ProcessBatch method with the test jobs
	results := processor.ProcessBatch(testJobs)

	// Verify the length of the results matches the length of the input jobs
	if len(results) != len(testJobs) {
		t.Errorf("Expected %d results, but got %d", len(testJobs), len(results))
	}

	// Verify each result corresponds to the correct job
	for i, result := range results {
		// Check if the job ID in the result matches the ID of the corresponding input job
		if len(result.JobIDs) != 1 || result.JobIDs[0] != testJobs[i].ID {
			t.Errorf("Expected result for job ID %d, but got %v", testJobs[i].ID, result.JobIDs)
		}

		// Check if the data in the result matches the data of the corresponding input job
		if len(result.Data) != 1 || result.Data[0] != testJobs[i].Data {
			t.Errorf("Expected data in result for job ID %d to be %s, but got %v", testJobs[i].ID, testJobs[i].Data, result.Data)
		}

		// Check if the Errors field is nil
		if result.Errors != nil {
			t.Errorf("Expected no errors for job ID %d, but got %v", testJobs[i].ID, result.Errors)
		}
	}
}
