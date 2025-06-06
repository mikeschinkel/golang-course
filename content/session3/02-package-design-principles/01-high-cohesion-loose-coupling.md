### 1. High Cohesion, Loose Coupling
```go
// Good: Main business logic in focused package
package acme

type Account struct {
    ID       string
    Number   string
    Balance  int64    // cents to avoid float precision issues
    Currency string
}

type Customer struct {
    ID    string
    Email string
    Name  string
}

type Service struct {
    store    Store
    parser   ChatParser
    k8s      K8sClient
}

// Define interfaces for dependencies
type Store interface {
    GetAccount(id string) (*Account, error)
    SaveAccount(*Account) error
    GetCustomer(id string) (*Customer, error)
}

type ChatParser interface {
    ParseTransferRequest(msg string) (*TransferRequest, error)
}

type K8sClient interface {
    GetPodStatus(namespace, name string) (string, error)
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
