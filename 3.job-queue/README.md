# Job Queue with Go

A simple job queue implementation using goroutines and channels in Go.

## Project Structure

```
job-queue/
├── main.go          # Application entry point
├── go.mod           # Go module definition
├── job/
│   └── job.go       # Job definition and processing logic
├── worker/
│   └── worker.go    # Worker implementation
└── queue/
    └── queue.go     # Job queue management
```

## Key Components

### Job (`job/job.go`)
- Defines what a job looks like
- Contains job processing logic
- Simple structure with ID, Name, and Payload

### Worker (`worker/worker.go`)
- Processes jobs from the job channel
- Runs in its own goroutine
- Can be started and stopped gracefully

### Queue (`queue/queue.go`)
- Manages multiple workers
- Distributes jobs to available workers
- Handles worker lifecycle

## How It Works

1. **Job Queue Creation**: Creates a specified number of workers and a buffered channel for jobs
2. **Worker Start**: Each worker runs in its own goroutine, waiting for jobs
3. **Job Distribution**: Jobs are sent to the channel, and workers pick them up automatically
4. **Processing**: Workers process jobs one at a time
5. **Graceful Shutdown**: All workers can be stopped cleanly

## Running the Code

1. Create the project directory structure
2. Copy all the files to their respective locations
3. Run the following commands:

```bash
cd job-queue
go mod tidy
go run main.go
```

## Key Go Concepts Used

- **Goroutines**: Lightweight threads for concurrent execution
- **Channels**: Communication between goroutines
- **Select Statement**: Non-blocking channel operations
- **WaitGroup**: Synchronization to wait for goroutines to complete
- **Buffered Channels**: Allows queuing jobs without blocking

## Sample Output

```
Starting job queue with 3 workers
Worker 1 started
Worker 2 started  
Worker 3 started
Adding job 1 to queue
Worker 1 received job 1
Processing job 1: Send Email
...
```

## Customization

- Change the number of workers in `main.go`
- Modify buffer size for the job channel
- Add different job types by extending the `Job` struct
- Implement different processing logic in the `Process()` method