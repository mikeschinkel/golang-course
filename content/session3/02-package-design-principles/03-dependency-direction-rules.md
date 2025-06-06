### 3. Dependency Direction Rules
```go
// ✅ Good: Inner domains don't depend on outer layers
// acme package (core domain) defines interfaces
package acme

type PaymentProcessor interface {
    ProcessPayment(amount int) error
}

// ✅ Good: Infrastructure implements domain interfaces
// payment package (infrastructure) implements acme interfaces
package payment

import "github.com/myapp/acme"

type StripeProcessor struct{}

func (s *StripeProcessor) ProcessPayment(amount int) error {
    // Stripe-specific implementation
    return nil
}

// ❌ Bad: Core domain depending on infrastructure
package acme

import "github.com/stripe/stripe-go"  // External dependency in core!

type Account struct {
    stripeClient *stripe.Client  // Infrastructure leaking into domain
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
