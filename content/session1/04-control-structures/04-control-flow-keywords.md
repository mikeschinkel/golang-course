# Control Flow Keywords

```go
// Break and continue
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue  // Skip even numbers
    }
    if i > 7 {
        break     // Exit loop when i > 7
    }
    fmt.Println(i)  // Prints 1, 3, 5, 7
}

// Labeled break/continue
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer  // Break from outer loop
        }
        fmt.Printf("(%d,%d) ", i, j)
    }
}

// Goto (use sparingly)
func example() error {
    file, err := os.Open("test.txt")
    if err != nil {
        goto cleanup
    }
    
    // Process file...
    
    if someError {
        err = errors.New("processing error")
        goto cleanup
    }
    
    file.Close()
    return nil
    
cleanup:
    if file != nil {
        file.Close()
    }
    return err
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
