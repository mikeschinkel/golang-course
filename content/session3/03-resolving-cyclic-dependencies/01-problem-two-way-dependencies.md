### Problem: Two-Way Dependencies
```go
// ❌ This creates a cycle
// acme/account.go
package acme

import "github.com/myapp/transaction"

type Account struct {
    ID           string
    Transactions []transaction.Transaction  // acme depends on transaction
}

// transaction/transaction.go
package transaction

import "github.com/myapp/acme"

type Transaction struct {
    ID      string
    Account acme.Account  // transaction depends on acme - CYCLE!
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
