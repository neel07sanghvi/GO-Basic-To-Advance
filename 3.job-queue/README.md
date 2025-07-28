# Job Queue with Graceful Shutdown

A job queue implementation using goroutines, channels, and context cancellation for graceful shutdown in Go.

## Project Structure

```
job-queue/
├── main.go          # Application entry point with signal handling
├── go.mod           # Go module definition
├── job/
│   └── job.go       # Job definition with context-aware processing
├── worker/
│   └── worker.go    # Worker implementation with context cancellation
└── queue/
    └── queue.go     # Job queue with graceful shutdown management
```

### What is Graceful Shutdown?
- **Stops accepting new jobs** when shutdown is initiated
- **Allows current jobs to complete** before terminating workers
- **Provides timeout mechanism** to prevent hanging
- **Handles system signals** (Ctrl+C, SIGTERM) properly

### What is Context Cancellation?
- **Context**: Go's standard way to carry cancellation signals
- **Propagation**: Cancellation signal spreads through all goroutines
- **Responsive**: Jobs can check if they should stop mid-processing
- **Clean**: No abrupt termination or resource leaks

## Key Components

### Enhanced Job Processing (`job/job.go`)
```go
// New method that respects context cancellation
func (j Job) ProcessWithContext(ctx context.Context) error {
    select {
    case <-time.After(workTime):  // Normal completion
        return nil
    case <-ctx.Done():           // Cancelled during processing
        return ctx.Err()
    }
}
```

### Context-Aware Workers (`worker/worker.go`)
```go
func (w *Worker) Start(ctx context.Context) {
    for {
        select {
        case job := <-w.JobChan:     // Process job
        case <-ctx.Done():           // Shutdown signal received
            return
        }
    }
}
```

### Graceful Queue Management (`queue/queue.go`)
- **Shutdown()**: Graceful shutdown with timeout
- **ForceShutdown()**: Immediate termination
- **Context management**: Coordinates all workers
- **Signal handling**: Responds to OS signals

## How Graceful Shutdown Works

### 1. Signal Detection
```go
// Catches Ctrl+C, SIGTERM, etc.
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
```

### 2. Shutdown Phases
```
Phase 1: Stop accepting new jobs
    ↓
Phase 2: Let current jobs finish (with timeout)
    ↓
Phase 3: If timeout exceeded, force cancellation
    ↓
Phase 4: Wait for all workers to exit
```

### 3. Context Propagation
```
Main Context (cancelled on shutdown)
    ├── Worker 1 Context ──→ Stops processing
    ├── Worker 2 Context ──→ Stops processing
    └── Worker 3 Context ──→ Stops processing
```

## Running the Code

```bash
cd job-queue
go mod tidy
go run main.go
```

### Testing Graceful Shutdown
1. Run the program: `go run main.go`
2. Press **Ctrl+C** while jobs are processing
3. Observe how current jobs finish before shutdown
4. Notice no abrupt termination

## Key Benefits

### Before (Basic Shutdown)
- ❌ Jobs might be interrupted mid-processing
- ❌ No timeout handling
- ❌ Potential resource leaks
- ❌ Poor user experience

### After (Graceful Shutdown)
- ✅ Current jobs complete safely
- ✅ Configurable shutdown timeout
- ✅ Clean resource cleanup
- ✅ Production-ready reliability

## Real-World Applications

### When This Matters
- **Payment processing**: Don't interrupt money transfers
- **File uploads**: Let uploads complete
- **Database operations**: Ensure data consistency
- **API requests**: Finish serving current requests

### Production Scenarios
- **Container orchestration** (Kubernetes, Docker)
- **Load balancer draining** (removing servers from rotation)
- **Application updates** (zero-downtime deployments)
- **System maintenance** (planned shutdowns)

## Configuration Options

```go
// Shutdown timeout - how long to wait for jobs to finish
err := jobQueue.Shutdown(5 * time.Second)

// Force shutdown - immediate termination
jobQueue.ForceShutdown()
```

## Output Example

```
Starting job queue with 3 workers
Adding job 1 to queue
Worker 1 received job 1
Processing job 1: Send Email
^C
Received signal: interrupt
Initiating graceful shutdown...
Job 1 completed: Welcome email to user@example.com
Worker 1 finished job 1
All workers finished gracefully
Program finished!
```

This implementation provides enterprise-grade reliability for job processing systems!