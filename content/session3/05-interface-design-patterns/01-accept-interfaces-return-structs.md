### 1. Accept Interfaces, Return Structs
```go
// ✅ Good
func NewUserService(store UserStore, notifier Notifier) *Service {
    return &Service{store: store, notifier: notifier}
}

// ❌ Avoid returning interfaces unless necessary
func NewUserService() UserService {  // Unnecessary abstraction
    return &Service{}
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
