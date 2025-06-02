# Select with Default
```go
select {
case msg := <-ch:
    // Handle message
default:
    // Non-blocking alternative
}
```