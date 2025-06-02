# Quick Reference - Common Patterns

```go
// String to int
num, err := strconv.Atoi("123")

// Int to string
str := strconv.Itoa(123)

// String formatting
msg := fmt.Sprintf("Hello %s, you have %d messages", name, count)

// Check if slice contains element
func contains(slice []int, item int) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// Remove element from slice
func remove(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}

// Copy slice
dest := make([]int, len(source))
copy(dest, source)

// Or
dest := append([]int(nil), source...)
```