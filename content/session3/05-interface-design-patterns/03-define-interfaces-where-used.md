### 3. Define Interfaces Where They're Used
```go
// ✅ Good: Interface defined in consumer package
// acme/service.go
package acme

type Store interface {  // Defined where it's used
    Save(*Account) error
    Get(id string) (*Account, error)
}

type Service struct {
    store Store
}

// postgres/account_store.go
package postgres

import "github.com/myapp/acme"

type AccountStore struct {
    db *sql.DB
}

// Implements acme.Store interface
func (s *AccountStore) Save(a *acme.Account) error { /* */ }
func (s *AccountStore) Get(id string) (*acme.Account, error) { /* */ }
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
