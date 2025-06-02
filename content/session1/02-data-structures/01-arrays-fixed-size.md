# Arrays (Fixed Size)

```go
// Array declaration
var arr [5]int                    // [0 0 0 0 0]
var arr2 = [5]int{1, 2, 3, 4, 5} // [1 2 3 4 5]
arr3 := [...]int{1, 2, 3}         // Size inferred: [3]int

// Array operations
len(arr)        // Length: 5
arr[0] = 10     // Assignment
x := arr[0]     // Access

// Arrays are values (copied when assigned)
arr4 := arr2    // Copies entire array
arr4[0] = 999   // Doesn't affect arr2
```