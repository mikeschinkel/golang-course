### 1. Test Organization and Naming
```go
// acme/service_test.go
package acme_test  // External test package

import (
    "testing"
    "github.com/acmebank/acme"
)

// Test function naming: Test_<Function>_<Scenario>_<Expected>
func TestService_Transfer_ValidAccounts_Success(t *testing.T) {
    // Arrange
    mockStore := &MockStore{}
    service := acme.NewService(mockStore, nil, nil)
    
    // Act
    err := service.Transfer("acc1", "acc2", 1000)
    
    // Assert
    if err != nil {
        t.Errorf("expected no error, got %v", err)
    }
}

// Table-driven tests for multiple scenarios
func TestService_Transfer(t *testing.T) {
    tests := []struct {
        name          string
        fromAccount   string
        toAccount     string
        amount        int64
        setupMock     func(*MockStore)
        expectedError string
    }{
        {
            name:        "successful transfer",
            fromAccount: "acc1",
            toAccount:   "acc2", 
            amount:      1000,
            setupMock: func(m *MockStore) {
                m.accounts["acc1"] = &acme.Account{ID: "acc1", Balance: 5000}
                m.accounts["acc2"] = &acme.Account{ID: "acc2", Balance: 1000}
            },
            expectedError: "",
        },
        {
            name:          "insufficient funds",
            fromAccount:   "acc1",
            toAccount:     "acc2",
            amount:        10000,
            setupMock: func(m *MockStore) {
                m.accounts["acc1"] = &acme.Account{ID: "acc1", Balance: 5000}
                m.accounts["acc2"] = &acme.Account{ID: "acc2", Balance: 1000}
            },
            expectedError: "insufficient funds",
        },
        {
            name:          "account not found",
            fromAccount:   "nonexistent",
            toAccount:     "acc2",
            amount:        1000,
            setupMock:     func(m *MockStore) {},
            expectedError: "account not found",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Arrange
            mockStore := &MockStore{accounts: make(map[string]*acme.Account)}
            tt.setupMock(mockStore)
            service := acme.NewService(mockStore, nil, nil)
            
            // Act
            err := service.Transfer(tt.fromAccount, tt.toAccount, tt.amount)
            
            // Assert
            if tt.expectedError == "" {
                if err != nil {
                    t.Errorf("expected no error, got %v", err)
                }
            } else {
                if err == nil {
                    t.Errorf("expected error containing %q, got nil", tt.expectedError)
                } else if !strings.Contains(err.Error(), tt.expectedError) {
                    t.Errorf("expected error containing %q, got %v", tt.expectedError, err)
                }
            }
        })
    }
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
