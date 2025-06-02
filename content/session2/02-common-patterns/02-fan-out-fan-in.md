# Fan-out, Fan-in
```go
func fanOut(input <-chan int) (<-chan int, <-chan int) {
    out1 := make(chan int)
    out2 := make(chan int)
    
    go func() {
        defer close(out1)
        defer close(out2)
        for val := range input {
            out1 <- val
            out2 <- val
        }
    }()
    
    return out1, out2
}
```