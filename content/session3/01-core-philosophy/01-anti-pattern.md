### ❌ Anti-Pattern: Technical Layers (MVC)
```
project/
├── controllers/
│   ├── user_controller.go
│   └── order_controller.go
├── models/
│   ├── user.go
│   └── order.go
├── services/
│   ├── user_service.go
│   └── order_service.go
└── repositories/
    ├── user_repository.go
    └── order_repository.go
```

**Problems:**
- High coupling between layers
- Inevitable cyclic dependencies
- Changes spread across multiple packages
- Not how Go standard library is organized
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
