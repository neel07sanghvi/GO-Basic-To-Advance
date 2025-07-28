package job

import (
	"context"
	"fmt"
	"time"
)

// Job represents a unit of work to be processed
type Job struct {
	ID      int
	Name    string
	Payload string
}

// Process simulates job processing (legacy method)
func (j *Job) Process() {
	fmt.Printf("Processing job %d: %s\n", j.ID, j.Name)

	// Simulate some work by sleeping
	time.Sleep(time.Millisecond * 500)

	fmt.Printf("Job %d completed: %s\n", j.ID, j.Payload)
}

// ProcessWithContext simulates job processing with context cancellation support
func (j *Job) ProcessWithContext(ctx context.Context) error {
	fmt.Printf("Processing job %d: %s\n", j.ID, j.Name)

	// Simulate work that can be cancelled
	select {
	case <-time.After(time.Millisecond * 500): // Simulate work duration
		fmt.Printf("Job %d completed: %s\n", j.ID, j.Payload)
		return nil

	case <-ctx.Done():
		// Context was cancelled (graceful shutdown triggered)
		fmt.Printf("Job %d cancelled during processing: %s\n", j.ID, j.Name)
		return ctx.Err()
	}
}
