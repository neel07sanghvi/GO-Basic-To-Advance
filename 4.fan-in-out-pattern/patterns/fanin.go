package patterns

import (
	"sync"

	"github.com/neel07sanghvi/fan-in-out-pattern/models"
)

// FanIn collects results from multiple worker channels into a single slice
func FanIn(workerChannels []<-chan models.Order) []models.Order {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var results []models.Order

	// Start a goroutine for each worker channel
	for _, ch := range workerChannels {
		wg.Add(1)

		go func(c <-chan models.Order) {
			defer wg.Done()

			// Collect all results from this worker
			for order := range c {
				mu.Lock()
				results = append(results, order)
				mu.Unlock()
			}
		}(ch)
	}

	wg.Wait()

	return results
}

// Alternative FanIn implementation using a single output channel
// This version streams results as they come in rather than collecting them all
func FanInStream(workerChannels []<-chan models.Order) <-chan models.Order {
	var wg sync.WaitGroup
	output := make(chan models.Order)

	// Start a goroutine for each input channel
	for _, ch := range workerChannels {
		wg.Add(1)

		go func(c <-chan models.Order) {
			defer wg.Done()

			// Forward all values from c to output
			for order := range c {
				output <- order
			}
		}(ch)
	}

	// Starts a goroutine to wait for all worker channels to finish, then closes the output channel.
	// This runs in the background so the function can return the output channel immediately,
	// allowing the caller to start reading data without waiting.
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
