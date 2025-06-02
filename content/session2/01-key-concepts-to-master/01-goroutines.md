# Goroutines
```go
// Basic goroutine
go fmt.Println("Hello from goroutine!")

// Goroutine with anonymous function
go func() {
    fmt.Println("Anonymous goroutine")
}()

// Goroutine with function call
go processData(data)
```

**Critical Point:** Main function doesn't wait for goroutines to finish!