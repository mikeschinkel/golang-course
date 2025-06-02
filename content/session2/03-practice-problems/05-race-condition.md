# Problem 5: Race Condition
Fix this race condition:
```go
var counter int

func main() {
    for i := 0; i < 1000; i++ {
        go func() {
            counter++ // Race condition!
        }()
    }
    time.Sleep(time.Second)
    fmt.Println(counter) // Unpredictable result
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
