# Pointers with Structs

```go
type Point struct {
    X, Y int
}

// Create struct pointer
p1 := &Point{X: 10, Y: 20}    // Struct literal with address
p2 := new(Point)              // Zero value struct
p3 := &Point{}                // Zero value struct

// Automatic dereferencing with structs
p1.X = 15              // Same as (*p1).X = 15
fmt.Println(p1.Y)      // Same as (*p1).Y

// Passing to functions
func movePoint(p *Point, dx, dy int) {
    p.X += dx          // Modifies original
    p.Y += dy
}

point := Point{X: 5, Y: 10}
movePoint(&point, 3, 4)
fmt.Println(point)     // {8 14}
```