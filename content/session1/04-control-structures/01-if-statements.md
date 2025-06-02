# If Statements

```go
// Basic if
if x > 10 {
    fmt.Println("x is greater than 10")
}

// If with else
if x%2 == 0 {
    fmt.Println("even")
} else {
    fmt.Println("odd")
}

// If with else if
if score >= 90 {
    grade = "A"
} else if score >= 80 {
    grade = "B"
} else if score >= 70 {
    grade = "C"
} else {
    grade = "F"
}

// If with initialization
if err := someFunction(); err != nil {
    return err
}
// err is scoped to if block

// Complex conditions
if x > 0 && y > 0 && x*y < 100 {
    fmt.Println("Valid coordinates")
}
```