# Maps (Hash Tables)

```go
// Map creation
var m map[string]int              // nil map (read-only)
m = make(map[string]int)          // Empty map
m = map[string]int{               // Map literal
    "apple":  5,
    "banana": 3,
    "orange": 8,
}

// Map operations
m["grape"] = 10                   // Set value
value := m["apple"]               // Get value (0 if not exists)
value, ok := m["apple"]           // Check existence
if ok {
    fmt.Println("Found apple:", value)
}

delete(m, "banana")               // Delete key
len(m)                           // Number of key-value pairs

// Iterate over map
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// Keys only
for key := range m {
    fmt.Println(key)
}

// Set Implementation (Lookup-Only Maps)
// NAIVE: Using boolean values (wastes memory)
seen := map[string]bool{
    "alice":   true,
    "bob":     true, 
    "charlie": true,
}

// IDIOMATIC: Using empty struct (zero memory for values)
seen := map[string]struct{}{
    "alice":   {},
    "bob":     {}, 
    "charlie": {},
}

// Why struct{} is better:
// - bool takes 1 byte per value
// - struct{} takes 0 bytes per value
// - For large sets, this saves significant memory

// Usage is the same:
if _, exists := seen["alice"]; exists {
    fmt.Println("Alice is in the set")
}

// Adding to set
seen["diana"] = struct{}{}

// Or more readable with a type alias:
type StringSet map[string]struct{}

func (s StringSet) Add(key string) {
    s[key] = struct{}{}
}

func (s StringSet) Contains(key string) bool {
    _, exists := s[key]
    return exists
}

func (s StringSet) Remove(key string) {
    delete(s, key)
}

// Usage
userSet := make(StringSet)
userSet.Add("alice")
userSet.Add("bob")

if userSet.Contains("alice") {
    fmt.Println("Found alice")
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
