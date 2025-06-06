### 6. Testing HTTP Handlers
```go
// cmd/server/handlers_test.go
package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    
    "github.com/acmebank/acme"
)

func TestTransferHandler(t *testing.T) {
    // Setup
    mockStore := &MockStore{
        accounts: map[string]*acme.Account{
            "acc1": {ID: "acc1", Balance: 10000},
            "acc2": {ID: "acc2", Balance: 5000},
        },
    }
    
    service := acme.NewService(mockStore, nil, nil)
    handler := NewTransferHandler(service)
    
    // Test data
    requestBody := map[string]interface{}{
        "from_account": "acc1",
        "to_account":   "acc2", 
        "amount":       1000,
    }
    
    body, _ := json.Marshal(requestBody)
    
    // Create request
    req := httptest.NewRequest("POST", "/transfer", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    
    // Record response
    rr := httptest.NewRecorder()
    
    // Execute
    handler.ServeHTTP(rr, req)
    
    // Assert
    if rr.Code != http.StatusOK {
        t.Errorf("expected status 200, got %d", rr.Code)
    }
    
    var response map[string]interface{}
    if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
        t.Fatalf("failed to unmarshal response: %v", err)
    }
    
    if response["status"] != "success" {
        t.Errorf("expected status success, got %v", response["status"])
    }
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
