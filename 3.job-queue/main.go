package main

import (
	"fmt"
	"time"

	"githum.com/neel07sanghvi/job-queue/job"
	"githum.com/neel07sanghvi/job-queue/queue"
)

func main() {

	jobQueue := queue.NewJobQueue(3, 10)

	jobQueue.Start()

	jobs := []job.Job{
		{ID: 1, Name: "Send Email", Payload: "Welcome email to user@example.com"},
		{ID: 2, Name: "Process Payment", Payload: "Payment of $100 for order #123"},
		{ID: 3, Name: "Generate Report", Payload: "Monthly sales report"},
		{ID: 4, Name: "Backup Data", Payload: "Database backup to cloud storage"},
		{ID: 5, Name: "Send SMS", Payload: "Order confirmation SMS"},
		{ID: 6, Name: "Update Inventory", Payload: "Update stock for product #456"},
		{ID: 7, Name: "Send Newsletter", Payload: "Weekly newsletter to subscribers"},
		{ID: 8, Name: "Clean Logs", Payload: "Remove old log files"},
	}

	for _, j := range jobs {
		jobQueue.AddJob(j)
		time.Sleep(time.Millisecond * 100) // Small delay between adding jobs
	}

	fmt.Println("\nLetting workers process jobs...")
	time.Sleep(time.Second * 5)

	jobQueue.Stop()

	fmt.Println("Program finished!")
}
