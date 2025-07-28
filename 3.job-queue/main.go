package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"githum.com/neel07sanghvi/job-queue/job"
	"githum.com/neel07sanghvi/job-queue/queue"
)

func main() {
	// Create a job queue with 3 workers and buffer size of 10
	jobQueue := queue.NewJobQueue(3, 10)

	// Start the job queue
	jobQueue.Start()

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create and add some jobs
	jobs := []job.Job{
		{ID: 1, Name: "Send Email", Payload: "Welcome email to user@example.com"},
		{ID: 2, Name: "Process Payment", Payload: "Payment of $100 for order #123"},
		{ID: 3, Name: "Generate Report", Payload: "Monthly sales report"},
		{ID: 4, Name: "Backup Data", Payload: "Database backup to cloud storage"},
		{ID: 5, Name: "Send SMS", Payload: "Order confirmation SMS"},
		{ID: 6, Name: "Update Inventory", Payload: "Update stock for product #456"},
		{ID: 7, Name: "Send Newsletter", Payload: "Weekly newsletter to subscribers"},
		{ID: 8, Name: "Clean Logs", Payload: "Remove old log files"},
		{ID: 9, Name: "Process Image", Payload: "Resize user profile image"},
		{ID: 10, Name: "Send Push Notification", Payload: "New message notification"},
	}

	// Start a goroutine to add jobs
	go func() {
		for _, j := range jobs {
			if err := jobQueue.AddJob(j); err != nil {
				fmt.Printf("Failed to add job %d: %v\n", j.ID, err)
				break
			}
			time.Sleep(time.Millisecond * 200) // Small delay between adding jobs
		}
	}()

	// Wait for interrupt signal or let it run for demo
	select {
	case sig := <-sigChan:
		fmt.Printf("\nReceived signal: %v\n", sig)
		fmt.Println("Initiating graceful shutdown...")

		// Attempt graceful shutdown with 5 second timeout
		if err := jobQueue.Shutdown(5 * time.Second); err != nil {
			fmt.Printf("Graceful shutdown failed: %v\n", err)
			fmt.Println("Performing force shutdown...")
			jobQueue.ForceShutdown()
		}

	case <-time.After(8 * time.Second):
		// Demo: shutdown after 8 seconds
		fmt.Println("\nDemo time expired, shutting down...")
		if err := jobQueue.Shutdown(3 * time.Second); err != nil {
			fmt.Printf("Graceful shutdown failed: %v\n", err)
			jobQueue.ForceShutdown()
		}
	}

	fmt.Println("Program finished!")
}
