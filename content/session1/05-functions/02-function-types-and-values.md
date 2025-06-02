# Function Types and Values

```go
// Function type
type Calculator func(int, int) int

// Functions as values
var operation Calculator = func(a, b int) int { return a + b }
result := operation(8, 12)

// Function literals (anonymous functions)
square := func(x int) int {
    return x * x
}

// Immediately invoked function
result := func(x, y int) int {
    return x * y
}(7, 6)

// Closures
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

triple := makeMultiplier(3)
fmt.Println(triple(4))  // 12
fmt.Println(triple(5))  // 15
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
