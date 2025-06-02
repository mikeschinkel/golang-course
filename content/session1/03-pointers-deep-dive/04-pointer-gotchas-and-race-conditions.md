# Pointer Gotchas and Race Conditions

```go
// Common mistake: pointer to loop variable
var addresses []*string
names := []string{"Alice", "Bob", "Charlie"}
for _, name := range names {
    addresses = append(addresses, &name)  // WRONG! All point to same variable
}
// All addresses point to final value ("Charlie")

// Correct approach
var addresses []*string
names := []string{"Alice", "Bob", "Charlie"}
for _, name := range names {
    nameCopy := name  // Create new variable
    addresses = append(addresses, &nameCopy)
}

// CRITICAL: Mixed receiver types and concurrency
type Inventory struct {
    mu    sync.Mutex
    items map[string]int
}

// Pointer receiver - thread safe
func (inv *Inventory) AddItem(name string, count int) {
    inv.mu.Lock()
    defer inv.mu.Unlock()
    inv.items[name] += count
}

// Value receiver - NOT thread safe! 
func (inv Inventory) GetTotal() int {
    inv.mu.Lock()         // Locks a COPY of the mutex!
    defer inv.mu.Unlock() // Unlocks the COPY!
    total := 0
    for _, count := range inv.items {
        total += count    // Race condition with AddItem()!
    }
    return total
}

// Correct: all pointer receivers
type SafeInventory struct {
    mu    sync.Mutex
    items map[string]int
}

func (inv *SafeInventory) AddItem(name string, count int) {
    inv.mu.Lock()
    defer inv.mu.Unlock()
    inv.items[name] += count
}

func (inv *SafeInventory) GetTotal() int {  // Pointer receiver!
    inv.mu.Lock()
    defer inv.mu.Unlock()
    total := 0
    for _, count := range inv.items {
        total += count
    }
    return total
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
