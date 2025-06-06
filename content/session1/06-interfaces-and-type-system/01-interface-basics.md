# Interface Basics

```go
// Interface definition
type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error)
}

// Interface composition
type ReadWriter interface {
    Reader
    Writer
}

// Empty interface
var value interface{}  // Can hold any type
//var value any        // Means the same as prior line
value = 42
value = "hello"
value = []int{1, 2, 3}
```

---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
