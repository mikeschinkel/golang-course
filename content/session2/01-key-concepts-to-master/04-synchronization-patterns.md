# Synchronization Patterns

## WaitGroup
```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        // Do work
        fmt.Printf("Worker %d done\n", id)
    }(i)
}

wg.Wait() // Wait for all goroutines
```

## Mutex
```go
var (
    counter int
    mu      sync.Mutex
)

func increment() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}
```

## Atomic Operations
```go
var counter int64

// Atomic increment
atomic.AddInt64(&counter, 1)

// Atomic load
value := atomic.LoadInt64(&counter)
```