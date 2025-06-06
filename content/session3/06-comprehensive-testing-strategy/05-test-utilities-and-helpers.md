### 5. Test Utilities and Helpers
```go
// acme/testutil_test.go
package acme_test

import (
    "testing"
    "github.com/acmebank/acme"
)

// Test data builders
func NewTestAccount(t *testing.T, opts ...func(*acme.Account)) *acme.Account {
    t.Helper()
    
    account := &acme.Account{
        ID:       "test-" + randomString(8),
        Number:   "ACC-" + randomString(6),
        Balance:  10000,
        Currency: "USD",
    }
    
    for _, opt := range opts {
        opt(account)
    }
    
    return account
}

func WithBalance(balance int64) func(*acme.Account) {
    return func(a *acme.Account) {
        a.Balance = balance
    }
}

func WithCurrency(currency string) func(*acme.Account) {
    return func(a *acme.Account) {
        a.Currency = currency
    }
}

// Usage in tests
func TestService_Transfer_DifferentCurrencies(t *testing.T) {
    usdAccount := NewTestAccount(t, WithBalance(10000), WithCurrency("USD"))
    eurAccount := NewTestAccount(t, WithBalance(10000), WithCurrency("EUR"))
    
    // Test logic...
}

// Assertion helpers
func AssertAccountBalance(t *testing.T, store acme.Store, accountID string, expectedBalance int64) {
    t.Helper()
    
    account, err := store.GetAccount(accountID)
    if err != nil {
        t.Fatalf("failed to get account %s: %v", accountID, err)
    }
    
    if account.Balance != expectedBalance {
        t.Errorf("account %s: expected balance %d, got %d", 
            accountID, expectedBalance, account.Balance)
    }
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
