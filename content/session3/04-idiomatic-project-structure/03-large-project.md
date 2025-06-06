### Large Project
```
myapp/
├── go.mod
├── cmd/                // Multiple binaries
│   ├── server/
│   │   └── main.go
│   ├── worker/
│   │   └── main.go
│   └── migrate/
│       └── main.go
├── internal/           // Private to this module
│   ├── user/
│   ├── order/
│   ├── payment/
│   └── config/
├── pkg/               // Public packages (if library)
│   └── client/
├── api/               // API definitions (OpenAPI, protobuf)
├── web/               // Static assets
├── docs/
└── scripts/
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
