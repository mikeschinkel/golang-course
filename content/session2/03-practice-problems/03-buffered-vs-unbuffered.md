# Problem 3: Buffered vs Unbuffered
```go
// What's the difference between these two?
ch1 := make(chan int)    // Unbuffered
ch2 := make(chan int, 1) // Buffered with capacity 1
```