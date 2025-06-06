### 2. Mock Implementation Patterns
```go
// acme/mock_test.go
package acme_test

import (
    "github.com/acmebank/acme"
)

// Mock Store implementation
type MockStore struct {
    accounts   map[string]*acme.Account
    customers  map[string]*acme.Customer
    
    // Track calls for verification
    getAccountCalls []string
    saveAccountCalls []*acme.Account
    
    // Control behavior
    shouldErrorOnGet  bool
    shouldErrorOnSave bool
}

func (m *MockStore) GetAccount(id string) (*acme.Account, error) {
    m.getAccountCalls = append(m.getAccountCalls, id)
    
    if m.shouldErrorOnGet {
        return nil, errors.New("database error")
    }
    
    account, exists := m.accounts[id]
    if !exists {
        return nil, acme.ErrAccountNotFound
    }
    
    // Return copy to prevent test interference
    copy := *account
    return &copy, nil
}

func (m *MockStore) SaveAccount(account *acme.Account) error {
    m.saveAccountCalls = append(m.saveAccountCalls, account)
    
    if m.shouldErrorOnSave {
        return errors.New("database error")
    }
    
    // Store copy
    copy := *account
    m.accounts[account.ID] = &copy
    return nil
}

// Verification helpers
func (m *MockStore) AssertGetAccountCalledWith(t *testing.T, expected string) {
    t.Helper()
    if len(m.getAccountCalls) == 0 {
        t.Fatal("expected GetAccount to be called, but it wasn't")
    }
    
    found := false
    for _, call := range m.getAccountCalls {
        if call == expected {
            found = true
            break
        }
    }
    
    if !found {
        t.Errorf("expected GetAccount to be called with %q, but it wasn't. Calls: %v", 
            expected, m.getAccountCalls)
    }
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
