# Defer, Panic, and Recover

```go
// Defer - executes when function returns
func readConfig(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Always executes when function returns
    
    // Read configuration...
    return nil
}

// Multiple defers (LIFO order)
func setupAndCleanup() {
    defer fmt.Println("Cleanup step 3")
    defer fmt.Println("Cleanup step 2") 
    defer fmt.Println("Cleanup step 1")
    fmt.Println("Setup complete")
}
// Output: Setup complete, Cleanup step 1, Cleanup step 2, Cleanup step 3

// Panic and recover
func safeOperation(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("operation failed: %v", r)
        }
    }()
    
    if b == 0 {
        panic("cannot divide by zero")
    }
    
    result = a / b
    return
}
```