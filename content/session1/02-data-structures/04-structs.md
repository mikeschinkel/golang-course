# Structs

```go
// Struct definition
type Person struct {
    Name    string
    Age     int
    Email   string
    birthday string  // Unexported field
}

// Struct creation
var p1 Person                     // Zero value
p2 := Person{}                    // Zero value
p3 := Person{                     // Named fields
    Name:  "Alice",
    Age:   30,
    Email: "alice@example.com",
}
p4 := Person{"Bob", 25, "bob@example.com", ""}  // Positional

// Anonymous structs
config := struct {
    Host string
    Port int
}{
    Host: "localhost",
    Port: 8080,
}

// Embedded structs (composition)
type Employee struct {
    Person          // Embedded struct
    EmployeeID int
    Department string
}

emp := Employee{
    Person: Person{Name: "Charlie", Age: 35},
    EmployeeID: 12345,
}

// Access embedded fields directly
fmt.Println(emp.Name)  // Same as emp.Person.Name
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
