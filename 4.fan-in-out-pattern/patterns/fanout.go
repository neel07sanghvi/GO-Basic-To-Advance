package patterns

import (
	"github.com/neel07sanghvi/fan-in-out-pattern/models"
	"github.com/neel07sanghvi/fan-in-out-pattern/workers"
)

// FanOut distributes orders to multiple worker goroutines
// Returns a slice of channels, each representing a worker's output
func FanOut(orders []models.Order, numWorkers int) []<-chan models.Order {
	// Create input channel for distributing work
	ordersChan := make(chan models.Order, len(orders))

	// Send all orders to the input channel
	go func() {
		defer close(ordersChan)

		for _, order := range orders {
			ordersChan <- order
		}
	}()

	// Create output channels for each worker
	workersChan := make([]<-chan models.Order, numWorkers)

	// Start workers
	for i := 0; i < numWorkers; i++ {
		workerOutput := make(chan models.Order)
		workersChan[i] = workerOutput

		// Start worker goroutine
		go func(workerID int, input <-chan models.Order, output chan<- models.Order) {
			defer close(output)

			// Process orders until input channel is closed
			for order := range input {
				processedOrder := workers.ProcessOrder(order, workerID)
				output <- processedOrder
			}
		}(i+1, ordersChan, workerOutput)
	}

	return workersChan
}
