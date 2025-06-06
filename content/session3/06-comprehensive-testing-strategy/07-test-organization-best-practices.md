### 7. Test Organization Best Practices

#### File Structure
```
acme/
├── account.go
├── account_test.go          // Unit tests for account
├── service.go  
├── service_test.go          // Unit tests for service
├── integration_test.go      // Integration tests
├── benchmark_test.go        // Benchmarks
├── mock_test.go            // Mock implementations
├── testutil_test.go        // Test utilities
└── example_test.go         // Example tests (for documentation)
```

#### Test Commands
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with detailed coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run only unit tests (fast)
go test -short ./...

# Run integration tests
go test -tags=integration ./...

# Run benchmarks
go test -bench=. ./...

# Run specific test
go test -run TestService_Transfer ./acme

# Verbose output
go test -v ./...

# Parallel execution
go test -parallel 8 ./...
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
