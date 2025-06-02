# Worker Pool
```go
func workerPool(jobs <-chan int, results chan<- int) {
    for job := range jobs {
        // Process job
        results <- job * 2
    }
}

// Usage
jobs := make(chan int, 100)
results := make(chan int, 100)

// Start workers
for w := 1; w <= 3; w++ {
    go workerPool(jobs, results)
}

// Send jobs
for j := 1; j <= 5; j++ {
    jobs <- j
}
close(jobs)

// Collect results
for r := 1; r <= 5; r++ {
    <-results
}
```