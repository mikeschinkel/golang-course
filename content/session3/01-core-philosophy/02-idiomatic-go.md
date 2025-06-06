### ✅ Idiomatic Go: Specific Domain Packages
```
acmebank/
├── go.mod
├── go.sum
├── README.md
├── acme/                    // Main business logic
│   ├── account.go           // Account domain
│   ├── transaction.go       // Transaction domain  
│   ├── customer.go          // Customer domain
│   ├── service.go           // Core business operations
│   ├── store.go             // Data access interfaces
│   └── acme_test.go         // Domain tests
├── internal/                // Private packages
│   ├── postgres/
│   │   ├── account_store.go
│   │   ├── migrations/
│   │   └── postgres_test.go
│   ├── k8sutils/           // Kubernetes utilities
│   │   ├── config.go
│   │   ├── health.go
│   │   └── k8sutils_test.go
│   └── chatparser/         // Chat message parsing
│       ├── parser.go
│       ├── validators.go
│       └── parser_test.go
├── cmd/
│   ├── server/
│   │   └── main.go
│   ├── migrator/
│   │   └── main.go
│   └── worker/
│       └── main.go
├── pkg/                    // Public packages (if needed)
│   └── acmeclient/         // Client library
│       ├── client.go
│       └── client_test.go
├── web/                    // Static assets
│   └── static/
├── deployments/            // K8s manifests
│   ├── base/
│   └── overlays/
└── scripts/
    ├── build.sh
    └── test.sh
```

**Benefits:**
- Specific, non-conflicting package names
- Clear separation of public (`pkg/`) vs private (`internal/`) code
- Business logic concentrated in main `acme` package
- Utility packages are specific and focused
- Proper module structure with `go.mod`
- Comprehensive testing strategy
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
