# Go Fan-in/Fan-out Pattern Demo

A simple Go project to learn concurrent programming patterns using an order processing example.

## What are Fan-in/Fan-out Patterns?

- **Fan-out**: Split work among multiple workers (like giving tasks to different employees)
- **Fan-in**: Collect results from multiple workers back into one place

## Project Structure

```
order-processor/
├── main.go              # Main program
├── models/order.go      # Order data structure  
├── workers/processor.go # Order processing logic
└── patterns/
    ├── fanout.go       # Distribute work to workers
    └── fanin.go        # Collect results from workers
```

## How to Run

1. Create the folder structure above
2. Copy the code from each file
3. Run these commands:

```bash
cd order-processor
go mod init order-processor
go run main.go
```

## What Happens

1. **Creates 6 sample orders** (Alice, Bob, Charlie, etc.)
2. **Fan-out**: Distributes orders to 3 workers running at the same time
3. **Processing**: Each worker takes 100ms-1s to process an order
4. **Fan-in**: Collects all processed orders back together
5. **Shows results**: Total revenue and processing time

## Sample Output

```
🚀 Starting Order Processing System
===================================
📦 Processing 6 orders using Fan-out/Fan-in patterns

📤 Fan-out: Distributing orders to workers...
📥 Fan-in: Collecting processed orders...
👷 Worker 1: Processing order #1 for Alice will take 941ms ⏱️
👷 Worker 3: Processing order #2 for Bob will take 127ms ⏱️
👷 Worker 2: Processing order #3 for Charlie will take 861ms ⏱️
✅ Worker 3: Order #2 completed successfully
👷 Worker 3: Processing order #4 for Diana will take 387ms ⏱️
✅ Worker 3: Order #4 completed successfully
👷 Worker 3: Processing order #5 for Eve will take 226ms ⏱️
✅ Worker 3: Order #5 completed successfully
👷 Worker 3: Processing order #6 for Frank will take 968ms ⏱️
✅ Worker 2: Order #3 completed successfully
✅ Worker 1: Order #1 completed successfully
❌ Worker 3: Order #6 failed to process

✅ All orders processed!
========================
Order #2: Bob - Status: Completed - Revenue: $800.00
Order #4: Diana - Status: Completed - Revenue: $600.00
Order #5: Eve - Status: Completed - Revenue: $300.00
Order #3: Charlie - Status: Completed - Revenue: $150.00
Order #1: Alice - Status: Completed - Revenue: $1200.00
Order #6: Frank - Status: Failed - Revenue: $80.00
...
💰 Total Revenue: $3050.00
⏱️  Processing Time: 1.710077039s

🎉 Order processing system completed successfully!
```

## Key Learning Points

- **Goroutines**: `go func()` runs code concurrently
- **Channels**: `chan` sends data between goroutines safely
- **WaitGroup**: `sync.WaitGroup` waits for all workers to finish
- **Mutex**: `sync.Mutex` prevents data races when multiple goroutines write

## Why This Matters

Without concurrency: 6 orders × 500ms average = **3 seconds**
With 3 workers: Limited by slowest order ≈ **1 second**

**Result: 3x faster processing!** 🚀