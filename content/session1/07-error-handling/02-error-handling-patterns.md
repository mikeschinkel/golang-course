# Error Handling Patterns

```go
// Standard error handling
data, err := fetchData()
if err != nil {
    return fmt.Errorf("failed to fetch data: %w", err)  // Wrap error
}

// Error checking in conditions
if err := validateRequest(request); err != nil {
    log.Printf("request validation failed: %v", err)
    return err
}

// Multiple error checks
conn, err := net.Dial("tcp", "localhost:8080")
if err != nil {
    return err
}
defer conn.Close()

data, err := io.ReadAll(conn)
if err != nil {
    return fmt.Errorf("reading from connection: %w", err)
}

// Error unwrapping (Go 1.13+)
if errors.Is(err, net.ErrClosed) {
    // Handle connection closed
}

var netErr *NetworkError
if errors.As(err, &netErr) {
    // Handle network error specifically
}
```