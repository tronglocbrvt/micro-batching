package cancellation

import (
	"testing"
)

// TestJobCancellation tests job cancellation mechanism.
func TestJobCancellation(t *testing.T) {
	// Initialize JobCancellation
	jobCancellation := NewJobCancellation()

	// Register cancellation function for a job
	jobID := 1
	cancelCalled := false
	jobCancellation.RegisterCancellation(jobID, func() {
		cancelCalled = true
	})

	// Cancel the job
	jobCancellation.Cancel(jobID)

	// Assert that the cancellation function was called
	if !cancelCalled {
		t.Errorf("Expected cancellation function to be called, but it wasn't")
	}
}
