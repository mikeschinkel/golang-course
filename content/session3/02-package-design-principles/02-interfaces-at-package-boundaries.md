### 2. Interfaces at Package Boundaries
```go
// acme/acme.go
package acme

// Define interface in consumer package
type Notifier interface {
    SendEmail(to, subject, body string) error
}

type Service struct {
    store    Store
    notifier Notifier  // Accept interface, not concrete type
}

// notification/notification.go
package notification

import "github.com/myapp/acme"

type EmailService struct {
    apiKey string
}

// Implement interface defined in acme package
func (e *EmailService) SendEmail(to, subject, body string) error {
    // Implementation
    return nil
}

// main.go
func main() {
    emailSvc := &notification.EmailService{}
    acmeSvc := acme.NewService(store, emailSvc)  // Dependency injection
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
