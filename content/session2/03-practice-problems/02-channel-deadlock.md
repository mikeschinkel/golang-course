# Problem 2: Channel Deadlock
```go
// Will this deadlock? Why?
func main() {
    ch := make(chan int)
    ch <- 42
    val := <-ch
    fmt.Println(val)
}
```