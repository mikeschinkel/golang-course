# Goroutine Variable Capture
```go
// WRONG
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i) // Prints 5 five times!
    }()
}

// CORRECT
for i := 0; i < 5; i++ {
    go func(id int) {
        fmt.Println(id) // Prints 0,1,2,3,4
    }(i)
}
```