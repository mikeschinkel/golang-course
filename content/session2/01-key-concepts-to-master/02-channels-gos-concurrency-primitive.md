# Channels - Go's Concurrency Primitive
```go
// Unbuffered channel (synchronous)
ch := make(chan int)

// Buffered channel (asynchronous up to buffer size)
ch := make(chan int, 5)

// Sending and receiving
ch <- 42        // Send
value := <-ch   // Receive

// Closing channels
close(ch)

// Check if channel is closed
value, ok := <-ch
if !ok {
    fmt.Println("Channel is closed")
}
```