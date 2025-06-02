# Variable Declarations

```go
// Four ways to declare variables
var x int                    // Zero value: 0
var y int = 42              // Explicit initialization
var z = 42                  // Type inference
w := 42                     // Short declaration (only in functions)

// Multiple declarations
var (
    name    string
    age     int
    height  float64
)

// Multiple assignment
a, b := 10, 20
x, y = y, x                 // Swap values

// Blank identifier
_, err := strconv.Atoi("123")  // Ignore first return value
```