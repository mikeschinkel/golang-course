### 3. Integration Tests
```go
// internal/postgres/integration_test.go
//go:build integration
// +build integration

package postgres_test

import (
    "database/sql"
    "os"
    "testing"
    
    "github.com/acmebank/acme"
    "github.com/acmebank/internal/postgres"
    _ "github.com/lib/pq"
)

func TestAccountStore_Integration(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test in short mode")
    }
    
    // Setup test database
    dbURL := os.Getenv("TEST_DATABASE_URL")
    if dbURL == "" {
        t.Skip("TEST_DATABASE_URL not set")
    }
    
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        t.Fatalf("failed to connect to test database: %v", err)
    }
    defer db.Close()
    
    // Clean and setup test data
    setupTestSchema(t, db)
    defer cleanupTestData(t, db)
    
    store := postgres.NewAccountStore(db)
    
    // Test round trip
    account := &acme.Account{
        ID:       "test-123",
        Number:   "ACC-001",
        Balance:  10000,
        Currency: "USD",
    }
    
    err = store.SaveAccount(account)
    if err != nil {
        t.Fatalf("failed to save account: %v", err)
    }
    
    retrieved, err := store.GetAccount("test-123")
    if err != nil {
        t.Fatalf("failed to get account: %v", err)
    }
    
    if retrieved.Balance != account.Balance {
        t.Errorf("expected balance %d, got %d", account.Balance, retrieved.Balance)
    }
}

func setupTestSchema(t *testing.T, db *sql.DB) {
    t.Helper()
    
    schema := `
    CREATE TABLE IF NOT EXISTS accounts (
        id VARCHAR PRIMARY KEY,
        account_number VARCHAR NOT NULL,
        balance BIGINT NOT NULL,
        currency VARCHAR(3) NOT NULL
    );
    `
    
    _, err := db.Exec(schema)
    if err != nil {
        t.Fatalf("failed to setup test schema: %v", err)
    }
}

func cleanupTestData(t *testing.T, db *sql.DB) {
    t.Helper()
    
    _, err := db.Exec("DELETE FROM accounts WHERE id LIKE 'test-%'")
    if err != nil {
        t.Logf("failed to cleanup test data: %v", err)
    }
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
