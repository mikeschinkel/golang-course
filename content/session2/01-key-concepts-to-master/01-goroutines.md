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
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
