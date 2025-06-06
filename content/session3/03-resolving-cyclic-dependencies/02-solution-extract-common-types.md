### Solution 1: Extract Common Types
```go
// types/types.go - shared types package
package types

type AccountID string
type TransactionID string

// acme/account.go
package acme

import "github.com/myapp/types"

type Account struct {
    ID   types.AccountID
    Name string
}

// transaction/transaction.go
package transaction

import "github.com/myapp/types"

type Transaction struct {
    ID        types.TransactionID
    AccountID types.AccountID  // Reference by ID, not full object
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
