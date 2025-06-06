### 2. Small, Focused Interfaces
```go
// ✅ Good: Small, single-purpose interfaces
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type Closer interface {
    Close() error
}

// Compose when needed
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// ❌ Bad: Large, do-everything interfaces
type UserManager interface {
    CreateUser(name string) error
    UpdateUser(id string, name string) error
    DeleteUser(id string) error
    SendEmail(to, subject, body string) error  // Wrong abstraction level
    ValidatePassword(password string) bool
    GenerateReport() ([]byte, error)
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
