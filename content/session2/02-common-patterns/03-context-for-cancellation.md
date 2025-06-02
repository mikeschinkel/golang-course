# Context for Cancellation
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

select {
case result := <-doWork():
    return result
case <-ctx.Done():
    return ctx.Err() // timeout or cancellation
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
