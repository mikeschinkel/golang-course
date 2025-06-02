# Zero Values

```go
var i int        // 0
var f float64    // 0.0
var b bool       // false
var s string     // ""
var p *int       // nil
var slice []int  // nil
var m map[string]int // nil
var ch chan int  // nil
var fn func()    // nil

// Structs get zero values for all fields
type Person struct {
    Name string
    Age  int
}
var person Person  // {Name: "", Age: 0}
```