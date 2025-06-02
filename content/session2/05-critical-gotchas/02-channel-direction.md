# Channel Direction
```go
// Send-only channel
func sender(ch chan<- int) {
    ch <- 42
}

// Receive-only channel
func receiver(ch <-chan int) {
    val := <-ch
}
```