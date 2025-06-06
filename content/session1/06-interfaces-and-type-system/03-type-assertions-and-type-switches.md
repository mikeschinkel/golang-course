# Type Assertions and Type Switches

```go
// Type assertion
var data any = "world"
text := data.(string)           // Direct assertion (panics if wrong)
text, ok := data.(string)       // Safe assertion
if ok {
    fmt.Println("Text:", text)
}

// Type switch
switch v := data.(type) {
case string:
    fmt.Printf("String: %s\n", v)
case int:
    fmt.Printf("Integer: %d\n", v)
case bool:
    fmt.Printf("Boolean: %t\n", v)
default:
    fmt.Printf("Unknown: %T\n", v)
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
