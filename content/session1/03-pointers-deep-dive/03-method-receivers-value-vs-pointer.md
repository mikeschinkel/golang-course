# Method Receivers - Value vs Pointer

```go
type Circle struct {
    Radius float64
}

// Value receiver - operates on copy
func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

// Pointer receiver - can modify original
func (c *Circle) Resize(factor float64) {
    c.Radius *= factor
}

// Both work on values and pointers (Go auto-converts)
circle := Circle{Radius: 5.0}
circle.Resize(2)        // Go converts to (&circle).Resize(2)
area := circle.Area()   // area = 314.159

circlePtr := &Circle{Radius: 3.0}
circlePtr.Resize(1.5)   // Direct pointer call
area2 := circlePtr.Area() // Go converts to (*circlePtr).Area()
```