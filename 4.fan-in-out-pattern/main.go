package main

import (
	"fmt"
	"time"

	"github.com/neel07sanghvi/fan-in-out-pattern/models"
	"github.com/neel07sanghvi/fan-in-out-pattern/patterns"
)

func main() {
	fmt.Println("🚀 Starting Order Processing System")
	fmt.Println("===================================")

	// Create sample orders
	orders := []models.Order{
		{ID: 1, CustomerName: "Alice", Items: []string{"Laptop", "Mouse"}, Total: 1200.00},
		{ID: 2, CustomerName: "Bob", Items: []string{"Phone"}, Total: 800.00},
		{ID: 3, CustomerName: "Charlie", Items: []string{"Headphones", "Cable"}, Total: 150.00},
		{ID: 4, CustomerName: "Diana", Items: []string{"Tablet", "Case", "Stylus"}, Total: 600.00},
		{ID: 5, CustomerName: "Eve", Items: []string{"Monitor"}, Total: 300.00},
		{ID: 6, CustomerName: "Frank", Items: []string{"Keyboard", "Mouse Pad"}, Total: 80.00},
	}

	fmt.Printf("📦 Processing %d orders using Fan-out/Fan-in patterns\n\n", len(orders))

	start := time.Now()

	// Step 1: Fan-out - Distribute orders to multiple workers
	fmt.Println("📤 Fan-out: Distributing orders to workers...")
	processedOrdersChan := patterns.FanOut(orders, 3) // 3 workers

	// Step 2: Fan-in - Collect results from all workers
	fmt.Println("📥 Fan-in: Collecting processed orders...")
	results := patterns.FanIn(processedOrdersChan)

	duration := time.Since(start)

	// Display results
	fmt.Println("\n✅ All orders processed!")
	fmt.Println("========================")

	totalRevenue := 0.0

	for _, result := range results {
		fmt.Printf("Order #%d: %s - Status: %s - Revenue: $%.2f\n",
			result.ID, result.CustomerName, result.Status, result.Total)

		if result.Status == "Completed" {
			totalRevenue += result.Total
		}
	}

	fmt.Printf("\n💰 Total Revenue: $%.2f\n", totalRevenue)
	fmt.Printf("⏱️  Processing Time: %v\n", duration)
	fmt.Println("\n🎉 Order processing system completed successfully!")
}
