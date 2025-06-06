### Solution 2: Interface Segregation
```go
// acme/account.go
package acme

// Define minimal interface for what acme package needs
type TransactionSummary interface {
    GetID() string
    GetAmount() int64
}

type Account struct {
    ID           string
    transactions []TransactionSummary  // Depend on interface, not concrete type
}

// transaction/transaction.go
package transaction

type Transaction struct {
    ID     string
    Amount int64
}

func (t *Transaction) GetID() string     { return t.ID }
func (t *Transaction) GetAmount() int64  { return t.Amount }

// Transaction automatically satisfies acme.TransactionSummary
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
