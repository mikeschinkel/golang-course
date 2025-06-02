# Problem 3: Buffered vs Unbuffered
```go
// What's the difference between these two?
ch1 := make(chan int)    // Unbuffered
ch2 := make(chan int, 1) // Buffered with capacity 1
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
