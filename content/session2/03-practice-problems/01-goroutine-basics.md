# Problem 1: Goroutine Basics
```go
// What will this print?
func main() {
    go fmt.Println("1")
    fmt.Println("2")
}
// A) 1\n2  B) 2\n1  C) 2  D) 1
```