# Practice Problems - Rapid Fire

### Problem 1: Variables and Types
```go
// What's the output?
var x = 42
var y = &x
var z = &y

*y = 100
fmt.Println(x, *y, **z)
```

### Problem 2: Slices vs Arrays
```go
// What's the difference?
arr := [3]int{1, 2, 3}
slice := []int{1, 2, 3}

func modifyArray(a [3]int) { a[0] = 999 }
func modifySlice(s []int) { s[0] = 999 }

modifyArray(arr)
modifySlice(slice)
// What are arr and slice now?
```

### Problem 3: Map Operations
```go
// Fill in the blanks
m := make(map[string]int)
m["apple"] = 5

// Check if key exists and get value
value, _____ := m["banana"]
if _____ {
    fmt.Println("Found banana")
} else {
    fmt.Println("No banana")
}

// Delete a key
_______(m, "apple")
```

### Problem 4: Control Flow
```go
// Rewrite using switch
func getGrade(score int) string {
    if score >= 90 {
        return "A"
    } else if score >= 80 {
        return "B"
    } else if score >= 70 {
        return "C"
    } else if score >= 60 {
        return "D"
    } else {
        return "F"
    }
}
```

### Problem 5: Pointers and Methods
```go
type Counter struct {
    count int
}

// Implement these methods:
// - Increment() that increases count by 1
// - GetCount() that returns current count
// - Reset() that sets count to 0

// Which methods need pointer receivers and why?
```