package queue

import (
	"context"
	"fmt"
	"sync"
	"time"

	"githum.com/neel07sanghvi/job-queue/job"
	"githum.com/neel07sanghvi/job-queue/worker"
)

type JobQueue struct {
	workers   []*worker.Worker
	jobChan   chan job.Job
	wg        sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
	isRunning bool
	mu        sync.RWMutex // Protects isRunning
}

func NewJobQueue(numWorkers, bufferSize int) *JobQueue {
	jobChan := make(chan job.Job, bufferSize)
	workers := make([]*worker.Worker, numWorkers)

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < numWorkers; i++ {
		workers[i] = worker.NewWorker(i+1, jobChan)
	}

	return &JobQueue{
		workers: workers,
		jobChan: jobChan,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (jq *JobQueue) Start() {
	jq.mu.Lock()
	defer jq.mu.Unlock()

	fmt.Printf("Starting job queue with %d workers\n", len(jq.workers))
	jq.isRunning = true

	// Start all workers in separate goroutines
	for _, w := range jq.workers {
		jq.wg.Add(1)

		go func(worker *worker.Worker) {
			defer jq.wg.Done()
			worker.Start(jq.ctx)
		}(w)
	}
}

func (jq *JobQueue) AddJob(j job.Job) error {
	jq.mu.RLock()
	running := jq.isRunning
	jq.mu.RUnlock()

	if !running {
		return fmt.Errorf("queue is not running or is shutting down")
	}

	select {
	case jq.jobChan <- j:
		fmt.Printf("Adding job %d to queue\n", j.ID)
		return nil
	case <-jq.ctx.Done():
		return fmt.Errorf("queue is shutting down, cannot add job %d", j.ID)
	}
}

// Shutdown gracefully stops all workers
func (jq *JobQueue) Shutdown(timeout time.Duration) error {
	jq.mu.Lock()

	if !jq.isRunning {
		jq.mu.Unlock()
		return fmt.Errorf("queue is not running")
	}

	jq.isRunning = false
	jq.mu.Unlock()

	fmt.Println("Initiating graceful shutdown...")

	// Close job channel to prevent new jobs and signal workers
	close(jq.jobChan)

	// Create a timeout context for shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), timeout)
	defer shutdownCancel()

	// Channel to signal when all workers are done
	done := make(chan struct{})
	go func() {
		jq.wg.Wait() // Wait for all workers to finish
		close(done)
	}()

	// Wait for either all workers to finish or timeout
	select {
	case <-done:
		fmt.Println("All workers finished gracefully")

		return nil
	case <-shutdownCtx.Done():
		fmt.Println("Shutdown timeout reached, forcing cancellation...")
		jq.cancel()  // Cancel context to force workers to stop
		jq.wg.Wait() // Wait for forced shutdown

		return fmt.Errorf("shutdown timeout exceeded")
	}
}

// ForceShutdown immediately cancels all workers
func (jq *JobQueue) ForceShutdown() {
	jq.mu.Lock()
	jq.isRunning = false
	jq.mu.Unlock()

	fmt.Println("Force shutdown initiated...")
	jq.cancel() // Cancel context immediately
	close(jq.jobChan)
	jq.wg.Wait()
	fmt.Println("Force shutdown completed")
}
