# Function Basics

```go
// Basic function
func add(x, y int) int {
    return x + y
}

// Multiple parameters of same type
func multiply(x, y, z int) int {
    return x * y * z
}

// Multiple return values
func divmod(a, b int) (int, int) {
    return a / b, a % b
}

// Named return values
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = errors.New("division by zero")
        return  // Returns zero value for result and the error
    }
    result = a / b
    return  // Returns result and nil for err
}

// Variadic functions
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage
result := sum(1, 2, 3, 4, 5)
slice := []int{1, 2, 3}
result2 := sum(slice...)  // Expand slice
```