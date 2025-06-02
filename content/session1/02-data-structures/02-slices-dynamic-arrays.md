# Slices (Dynamic Arrays)

```go
// Slice creation
var slice []int                    // nil slice
slice = make([]int, 5)            // [0 0 0 0 0], length 5
slice = make([]int, 3, 5)         // length 3, capacity 5
slice = []int{1, 2, 3, 4, 5}      // slice literal

// Slice operations
len(slice)                        // Length
cap(slice)                        // Capacity
slice = append(slice, 6)          // Add element
slice = append(slice, 7, 8, 9)    // Add multiple

// Slicing
s := []int{0, 1, 2, 3, 4, 5}
s[1:4]     // [1 2 3] - from index 1 to 3
s[:3]      // [0 1 2] - from start to index 2
s[2:]      // [2 3 4 5] - from index 2 to end
s[:]       // [0 1 2 3 4 5] - entire slice

// Copy slices
dest := make([]int, len(slice))
copy(dest, slice)
```