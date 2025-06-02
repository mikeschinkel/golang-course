# Error Interface

```go
// Error interface
type error interface {
    Error() string
}

// Creating errors
err1 := errors.New("network connection failed")
err2 := fmt.Errorf("user %s not authorized", username)

// Custom error types
type NetworkError struct {
    Host    string
    Port    int
    Message string
}

func (e *NetworkError) Error() string {
    return fmt.Sprintf("network error connecting to %s:%d - %s", e.Host, e.Port, e.Message)
}

// Usage
func connectToServer(host string, port int) error {
    if port < 1 || port > 65535 {
        return &NetworkError{
            Host:    host,
            Port:    port,
            Message: "invalid port number",
        }
    }
    return nil
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
