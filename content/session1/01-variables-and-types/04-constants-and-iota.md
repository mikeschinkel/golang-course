# Constants and iota

```go
// Constants
const Pi = 3.14159
const (
    StatusOK     = 200
    StatusForbidden = 403
    StatusNotFound  = 404
)

// iota - automatic incrementing
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)

// iota with expressions
const (
    Regular = 100 * iota  // 0 * 100 = 0
    Premium               // 1 * 100 = 100  
    VIP                   // 2 * 100 = 200
    Enterprise            // 3 * 100 = 300
)

// Reset iota in new const block
const (
    Red = iota    // 0
    Green         // 1
    Blue          // 2
)
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
