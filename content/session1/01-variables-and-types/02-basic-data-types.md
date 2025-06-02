# Basic Data Types

```go
// Numeric types
var i int = 42                    // Platform dependent (32 or 64 bit)
var i8 int8 = 127                // 8-bit signed (-128 to 127)
var i16 int16 = 32767            // 16-bit signed
var i32 int32 = 2147483647       // 32-bit signed (rune alias)
var i64 int64 = 9223372036854775807

var ui uint = 42                 // Platform dependent unsigned
var ui8 uint8 = 255             // 8-bit unsigned (byte alias)
var ui16 uint16 = 65535         // 16-bit unsigned
var ui32 uint32 = 4294967295    // 32-bit unsigned
var ui64 uint64 = 18446744073709551615

var f32 float32 = 3.14          // 32-bit floating point
var f64 float64 = 3.14159265359 // 64-bit floating point (default)

var c64 complex64 = 1 + 2i      // Complex numbers
var c128 complex128 = 1 + 2i

// String and boolean
var s string = "Hello, 世界"     // UTF-8 encoded
var b bool = true               // true or false

// Type declarations
type UserID int64
type EmailAddress string

// Type aliass
type UserID=int64
type EmailAddress=string
```