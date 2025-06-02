# Implementing Interfaces

```go
type Logger struct {
    prefix string
}

// Implement Writer interface (implicit)
func (l *Logger) Write(data []byte) (int, error) {
    message := l.prefix + string(data)
    fmt.Print(message)
    return len(data), nil
}

// Usage
var w Writer = &Logger{prefix: "[INFO] "}
w.Write([]byte("Application started"))
```