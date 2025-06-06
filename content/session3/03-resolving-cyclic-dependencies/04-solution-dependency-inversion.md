### Solution 3: Dependency Inversion
```go
// Instead of direct dependencies, use a mediator pattern
// events/events.go
package events

type AccountCreated struct {
    AccountID string
    CustomerID string
}

type TransactionProcessed struct {
    TransactionID string
    AccountID     string
    Amount        int64
}

// acme/service.go
package acme

import "github.com/myapp/events"

type Service struct {
    eventBus events.Publisher
}

func (s *Service) CreateAccount(customerID string) error {
    account := &Account{CustomerID: customerID}
    // Save account...
    
    // Publish event instead of direct call
    return s.eventBus.Publish(events.AccountCreated{
        AccountID:  account.ID,
        CustomerID: account.CustomerID,
    })
}

// transaction/service.go
package transaction

import "github.com/myapp/events"

type Service struct {
    eventBus events.Subscriber
}

func (s *Service) HandleAccountCreated(event events.AccountCreated) error {
    // Handle account creation in transaction context
    return nil
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
