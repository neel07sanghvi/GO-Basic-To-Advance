package queue

import (
	"fmt"
	"sync"

	"githum.com/neel07sanghvi/job-queue/job"
	"githum.com/neel07sanghvi/job-queue/worker"
)

type JobQueue struct {
	workers   []*worker.Worker
	jobChan   chan job.Job
	wg        sync.WaitGroup
	isRunning bool
}

func NewJobQueue(numWorkers, bufferSize int) *JobQueue {
	jobChan := make(chan job.Job, bufferSize)
	workers := make([]*worker.Worker, numWorkers)

	for i := 0; i < numWorkers; i++ {
		workers[i] = worker.NewWorker(i+1, jobChan)
	}

	return &JobQueue{
		workers: workers,
		jobChan: jobChan,
	}
}

func (jq *JobQueue) Start() {
	fmt.Printf("Starting job queue with %d workers\n", len(jq.workers))
	jq.isRunning = true

	// Start all workers in separate goroutines
	for _, w := range jq.workers {
		jq.wg.Add(1)

		go func(worker *worker.Worker) {
			defer jq.wg.Done()
			worker.Start()
		}(w)
	}
}

func (jq *JobQueue) AddJob(j job.Job) {
	if !jq.isRunning {
		fmt.Println("Queue is not running!")
		return
	}

	fmt.Printf("Adding job %d to queue\n", j.ID)
	jq.jobChan <- j
}

func (jq *JobQueue) Stop() {
	fmt.Println("Stopping job queue...")
	jq.isRunning = false

	close(jq.jobChan)

	for _, w := range jq.workers {
		w.Stop()
	}

	jq.wg.Wait()
	fmt.Println("Job queue stopped")
}
