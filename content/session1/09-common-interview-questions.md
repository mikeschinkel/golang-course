# Common Interview Questions

## 1. Memory and Pointers
**Q: "What's the zero value of a pointer in Go?"**

**Q: "When should you use pointer receivers vs value receivers?"**

**Q: "What's a dangerous pattern with mixed pointer and value receivers in concurrent code?"**

**Q: "Why can mixing receiver types create race conditions?"**

**Q: "What's wrong with having a mutex in a struct with value receivers?"**

**Q: "Explain method sets and how they relate to interface satisfaction."**

**Q: "What happens if you try to dereference a nil pointer?"**

**Q: "Explain the difference between `new()` and `make()` in Go."**

**Q: "How does Go handle memory management?"**

## 2. Data Types and Structures
**Q: "What's the difference between arrays and slices in Go?"**

**Q: "How do you check if a key exists in a map?"**

**Q: "What happens when you access a non-existent map key?"**

**Q: "Explain how the `append()` function works with slices."**

**Q: "What are the zero values for different Go types?"**

**Q: "What's the most memory-efficient way to implement a set in Go?"**

**Q: "Why would you use `map[string]struct{}` instead of `map[string]bool`?"**

**Q: "What's the size of an empty struct in Go?"**

## 3. Control Flow and Functions
**Q: "Go only has one loop construct. What is it and how is it used?"**

**Q: "What's the difference between `break` and `continue` in Go?"**

**Q: "Explain how `defer` works and when it executes."**

**Q: "What happens if you have multiple `defer` statements?"**

**Q: "How do you handle multiple return values in Go?"**

## 4. Interface and Type System
**Q: "How do you implement an interface in Go?"**

**Q: "What is the empty interface and when would you use it?"**

**Q: "Explain type assertions in Go."**

**Q: "What's the difference between type assertion and type conversion?"**

## 5. Error Handling
**Q: "How does error handling work in Go?"**

**Q: "How do you create custom error types?"**

**Q: "What's the idiomatic way to handle errors in Go?"**

**Q: "Explain error wrapping in Go 1.13+."**

## 6. Code Analysis Questions
**Q: "What's wrong with this code?"**
```go
func getUsers() []*User {
    var users []*User
    for i := 0; i < 3; i++ {
        user := &User{ID: i}
        users = append(users, &user.ID) // What's the issue?
    }
    return users
}
```

**Q: "Will this code compile? If not, why?"**
```go
func main() {
    var x int
    var p *int = x
    fmt.Println(*p)
}
```

**Q: "What will this print and why?"**
```go
func main() {
    slice := []int{1, 2, 3}
    for i, v := range slice {
        slice = append(slice, v+10)
        if i >= 2 {
            break
        }
    }
    fmt.Println(slice)
}
```