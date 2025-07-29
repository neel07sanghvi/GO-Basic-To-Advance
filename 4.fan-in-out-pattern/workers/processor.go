package workers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/neel07sanghvi/fan-in-out-pattern/models"
)

// ProcessOrder simulates order processing with random delays
func ProcessOrder(order models.Order, workerID int) models.Order {
	// Simulate processing time (100ms to 1s)
	processingTime := time.Duration(rand.Intn(900)+100) * time.Millisecond
	fmt.Printf("👷 Worker %d: Processing order #%d for %s will take %v ⏱️\n",
		workerID, order.ID, order.CustomerName, processingTime)
	time.Sleep(processingTime)

	order.Status = "Completed"
	order.ProcessedAt = time.Now()

	// Simulate occasional processing failures (10% chance)
	if rand.Float32() < 0.1 {
		order.Status = "Failed"
		fmt.Printf("❌ Worker %d: Order #%d failed to process\n", workerID, order.ID)
	} else {
		fmt.Printf("✅ Worker %d: Order #%d completed successfully\n", workerID, order.ID)
	}

	return order
}
