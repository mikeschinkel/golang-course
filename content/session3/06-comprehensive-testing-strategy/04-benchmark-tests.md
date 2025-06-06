### 4. Benchmark Tests
```go
// acme/benchmark_test.go
package acme_test

import (
    "testing"
    "github.com/acmebank/acme"
)

func BenchmarkService_Transfer(b *testing.B) {
    mockStore := &MockStore{
        accounts: map[string]*acme.Account{
            "acc1": {ID: "acc1", Balance: 1000000},
            "acc2": {ID: "acc2", Balance: 1000000},
        },
    }
    
    service := acme.NewService(mockStore, nil, nil)
    
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        // Alternate transfer direction to avoid running out of funds
        if i%2 == 0 {
            service.Transfer("acc1", "acc2", 100)
        } else {
            service.Transfer("acc2", "acc1", 100)
        }
    }
}

func BenchmarkChatParser_ParseTransferRequest(b *testing.B) {
    parser := chatparser.NewParser()
    message := "transfer $100 from account ACC-001 to account ACC-002"
    
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        _, err := parser.ParseTransferRequest(message)
        if err != nil {
            b.Fatal(err)
        }
    }
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
