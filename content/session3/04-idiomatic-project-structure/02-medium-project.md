### Medium Project
```
myapp/
├── go.mod
├── main.go
├── user/               // Domain packages
│   ├── user.go
│   ├── service.go
│   ├── store.go
│   └── user_test.go
├── order/
│   ├── order.go
│   ├── service.go
│   └── store.go
└── internal/           // Private packages
    └── config/
        └── config.go
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
