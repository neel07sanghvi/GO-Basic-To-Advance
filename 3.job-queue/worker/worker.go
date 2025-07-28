package worker

import (
	"context"
	"fmt"

	"githum.com/neel07sanghvi/job-queue/job"
)

type Worker struct {
	ID      int
	JobChan chan job.Job
}

func NewWorker(id int, jobChan chan job.Job) *Worker {
	return &Worker{
		ID:      id,
		JobChan: jobChan,
	}
}

func (w *Worker) Start(ctx context.Context) {
	fmt.Printf("Worder %d started\n", w.ID)

	for {
		select {
		case job, ok := <-w.JobChan:

			if !ok {
				fmt.Printf("Worker %d: job channel closed, stopping\n", w.ID)
				return
			}

			fmt.Printf("Worker %d received job %d\n", w.ID, job.ID)

			// Process job with context (allowing cancellation during processing)
			if err := job.ProcessWithContext(ctx); err != nil {
				fmt.Printf("Worker %d: job %d cancelled or failed: %v\n", w.ID, job.ID, err)
				continue
			}

			fmt.Printf("Worker %d finished job %d\n", w.ID, job.ID)

		case <-ctx.Done():
			fmt.Printf("Worker %d stopping due to context cancellation: %v\n", w.ID, ctx.Err())
			return
		}
	}
}
