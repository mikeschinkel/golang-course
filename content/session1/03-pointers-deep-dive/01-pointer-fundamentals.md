# Pointer Fundamentals

```go
// Pointer basics
var x int = 42
var p *int = &x      // p points to x
fmt.Println(*p)      // Dereference: prints 42

*p = 100             // Change value through pointer
fmt.Println(x)       // x is now 100

// Pointer zero value
var ptr *int         // nil
if ptr == nil {
    fmt.Println("Pointer is nil")
}

// Create pointer to new zero value
ptr = new(int)       // Points to new int with value 0
*ptr = 42
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
