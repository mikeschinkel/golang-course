# Go Interview Preparation - Answer Key
**FOR INSTRUCTOR USE ONLY**

## Session 1: Fundamentals & Memory Management

### Practice Problem Answers

#### Problem 1: Variables and Types
```go
var x = 42
var y = &x
var z = &y

*y = 100
fmt.Println(x, *y, **z)
```
**Answer:** `100 100 100`
- `x` is changed to 100 through pointer `y`
- All three expressions refer to the same value

#### Problem 2: Slices vs Arrays
```go
arr := [3]int{1, 2, 3}
slice := []int{1, 2, 3}

func modifyArray(a [3]int) { a[0] = 999 }
func modifySlice(s []int) { s[0] = 999 }

modifyArray(arr)
modifySlice(slice)
```
**Answer:** 
- `arr` remains `[1, 2, 3]` (arrays are passed by value)
- `slice` becomes `[999, 2, 3]` (slices are reference types)

#### Problem 3: Map Operations
```go
value, ok := m["banana"]
if ok {
    fmt.Println("Found banana")
} else {
    fmt.Println("No banana")
}

delete(m, "apple")
```

#### Problem 4: Control Flow
```go
func getGrade(score int) string {
    switch {
    case score >= 90:
        return "A"
    case score >= 80:
        return "B"
    case score >= 70:
        return "C"
    case score >= 60:
        return "D"
    default:
        return "F"
    }
}
```

#### Problem 5: Pointers and Methods
```go
func (c *Counter) Increment() {  // Pointer receiver - modifies original
    c.count++
}

func (c Counter) GetCount() int {  // Value receiver - read-only
    return c.count
}

func (c *Counter) Reset() {  // Pointer receiver - modifies original
    c.count = 0
}
```

### Interview Question Answers

#### Memory and Pointers
**Q: "What's the zero value of a pointer in Go?"**
**A:** `nil`

**Q: "When should you use pointer receivers vs value receivers?"**
**A:** 
- Pointer receivers: when you need to modify the receiver, for large structs (performance), or for consistency when some methods need pointer receivers
- Value receivers: for small structs, when you don't need to modify, for basic types

**Q: "What's a dangerous pattern with mixed pointer and value receivers in concurrent code?"**
**A:** Value receivers create copies of structs, including copies of mutexes. Locking a copied mutex provides no synchronization with the original, creating race conditions.

**Q: "Why can mixing receiver types create race conditions?"**
**A:** Because value receivers operate on copies while pointer receivers operate on originals. If a struct contains synchronization primitives (like mutexes), the value receiver will lock/unlock a copy, not the original mutex.

**Q: "What's wrong with having a mutex in a struct with value receivers?"**
**A:** The mutex gets copied with the struct. Locking a copied mutex doesn't synchronize with the original mutex, making the synchronization useless and creating race conditions.

**Q: "Explain method sets and how they relate to interface satisfaction."**
**A:** 
- Type T has methods with receiver T
- Type *T has methods with receiver T and *T  
- If interface methods have pointer receivers, only *T satisfies the interface
- If interface methods have value receivers, both T and *T satisfy the interface

**Q: "What happens if you try to dereference a nil pointer?"**
**A:** Runtime panic: "runtime error: invalid memory address or nil pointer dereference"

**Q: "Explain the difference between `new()` and `make()` in Go."**
**A:** 
- `new(T)` allocates zeroed storage for type T and returns `*T`
- `make()` is only for slices, maps, and channels - creates and initializes them

**Q: "How does Go handle memory management?"**
**A:** Garbage collection with escape analysis determining stack vs heap allocation

#### Data Types and Structures
**Q: "What's the difference between arrays and slices in Go?"**
**A:** 
- Arrays: fixed size, value types (copied when assigned)
- Slices: dynamic size, reference types (header points to underlying array)

**Q: "How do you check if a key exists in a map?"**
**A:** `value, ok := m[key]` - the second return value indicates existence

**Q: "What happens when you access a non-existent map key?"**
**A:** Returns the zero value for the value type, no panic

**Q: "Explain how the `append()` function works with slices."**
**A:** Adds elements to end of slice, may allocate new underlying array if capacity exceeded

**Q: "What are the zero values for different Go types?"**
**A:** int: 0, string: "", bool: false, pointer: nil, slice: nil, map: nil, interface: nil

**Q: "What's the most memory-efficient way to implement a set in Go?"**
**A:** `map[KeyType]struct{}` because `struct{}` has zero size, unlike `bool` which takes 1 byte per value

**Q: "Why would you use `map[string]struct{}` instead of `map[string]bool`?"**
**A:** `struct{}` has zero memory footprint while `bool` takes 1 byte. For large sets, this saves significant memory.

**Q: "What's the size of an empty struct in Go?"**
**A:** Zero bytes. `struct{}` is Go's zero-size type, making it perfect for set implementations.

#### Control Flow and Functions
**Q: "Go only has one loop construct. What is it and how is it used?"**
**A:** `for` loop with three variants:
- `for i := 0; i < 10; i++` (traditional)
- `for condition` (while-style)
- `for` (infinite)
- `for range` (iteration)

**Q: "What's the difference between `break` and `continue` in Go?"**
**A:** `break` exits the loop entirely, `continue` skips to next iteration

**Q: "Explain how `defer` works and when it executes."**
**A:** Executes when surrounding function returns, LIFO order for multiple defers

**Q: "What happens if you have multiple `defer` statements?"**
**A:** They execute in Last In, First Out (LIFO) order

**Q: "How do you handle multiple return values in Go?"**
**A:** Use multiple assignment: `result, err := function()` or blank identifier to ignore: `result, _ := function()`

#### Interface and Type System
**Q: "How do you implement an interface in Go?"**
**A:** Implicitly - any type that has the required methods automatically implements the interface

**Q: "What is the empty interface and when would you use it?"**
**A:** `interface{}` can hold any type, used for generic programming before generics

**Q: "Explain type assertions in Go."**
**A:** Extract concrete type from interface: `value, ok := interfaceVar.(ConcreteType)`

**Q: "What's the difference between type assertion and type conversion?"**
**A:** Type assertion extracts type from interface, type conversion changes one type to another

#### Error Handling
**Q: "How does error handling work in Go?"**
**A:** Functions return error as last value, check `if err != nil`

**Q: "How do you create custom error types?"**
**A:** Implement the `Error() string` method on your type

**Q: "What's the idiomatic way to handle errors in Go?"**
**A:** Check every error, handle or wrap and return up the call stack

**Q: "Explain error wrapping in Go 1.13+."**
**A:** Use `fmt.Errorf` with `%w` verb to wrap errors, then `errors.Is` and `errors.As` to unwrap

#### Code Analysis Answers
**Q: "What's wrong with this code?"**
```go
users = append(users, &user.ID) // Should be: users = append(users, user)
```
**Issue:** Taking address of field instead of the struct

**Q: "Will this code compile?"**
```go
var p *int = x  // Should be: var p *int = &x
```
**Issue:** Cannot assign int to *int, need address operator

**Q: "What will this print?"**
```go
slice := []int{1, 2, 3}
for i, v := range slice {
    slice = append(slice, v+10)
    if i >= 2 {
        break
    }
}
fmt.Println(slice)
```
**Answer:** `[1 2 3 11 12 13]`
**Explanation:** Range iterates over original slice length, appends don't affect current iteration

### Homework Assignment Solutions

#### Assignment 1: Data Types Mastery
```go
package main

import "fmt"

const (
    Read = 1 << (2 * iota)  // 4
    Write                   // 2  
    Execute                 // 1
)

func main() {
    // Four declaration methods
    var x int               // Zero value
    var y int = 42         // Explicit
    var z = 42             // Inference
    w := 42                // Short

    // Zero values demonstration
    var i int              // 0
    var f float64         // 0.0
    var b bool            // false
    var s string          // ""
    var p *int            // nil

    // Type conversions
    num := 42
    f64 := float64(num)
    str := fmt.Sprintf("%d", num)
}
```

#### Assignment 2: Collection Operations
```go
func reverseSlice(slice []int) []int {
    result := make([]int, len(slice))
    for i, v := range slice {
        result[len(slice)-1-i] = v
    }
    return result
}

func mergeMaps(m1, m2 map[string]int) map[string]int {
    result := make(map[string]int)
    for k, v := range m1 {
        result[k] = v
    }
    for k, v := range m2 {
        result[k] = v  // Overwrites
    }
    return result
}

func findDuplicates(strs []string) []string {
    seen := make(map[string]bool)
    dups := make(map[string]bool)
    var result []string
    
    for _, s := range strs {
        if seen[s] {
            if !dups[s] {
                result = append(result, s)
                dups[s] = true
            }
        } else {
            seen[s] = true
        }
    }
    return result
}

func groupByLength(strs []string) map[int][]string {
    result := make(map[int][]string)
    for _, s := range strs {
        length := len(s)
        result[length] = append(result[length], s)
    }
    return result
}
```

#### Assignment 3: Pointer Practice
```go
func swap(a, b *int) {
    *a, *b = *b, *a
}

type Node struct {
    Value int
    Next  *Node
}

type LinkedList struct {
    Head *Node
}

func (ll *LinkedList) Add(value int) {
    newNode := &Node{Value: value}
    if ll.Head == nil {
        ll.Head = newNode
    } else {
        current := ll.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
}

func (ll *LinkedList) Remove(value int) bool {
    if ll.Head == nil {
        return false
    }
    
    if ll.Head.Value == value {
        ll.Head = ll.Head.Next
        return true
    }
    
    current := ll.Head
    for current.Next != nil {
        if current.Next.Value == value {
            current.Next = current.Next.Next
            return true
        }
        current = current.Next
    }
    return false
}
```

#### Assignment 4: Error Handling
```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

func validateUser(name, email string) error {
    if name == "" {
        return &ValidationError{Field: "name", Message: "cannot be empty"}
    }
    if !strings.Contains(email, "@") {
        return &ValidationError{Field: "email", Message: "must contain @"}
    }
    return nil
}

func processUser(name, email string) error {
    if err := validateUser(name, email); err != nil {
        return fmt.Errorf("user validation failed: %w", err)
    }
    return nil
}

// Usage with error unwrapping
err := processUser("", "invalid")
var valErr *ValidationError
if errors.As(err, &valErr) {
    fmt.Printf("Validation failed on field: %s\n", valErr.Field)
}
```

#### Assignment 5: Mini Project - Bank Account
```go
package bank

import (
    "errors"
    "fmt"
)

type Account struct {
    ID       string
    Balance  int64  // Use int64 for money (cents)
    Currency string
}

func (a *Account) Deposit(amount int64) error {
    if amount <= 0 {
        return errors.New("deposit amount must be positive")
    }
    a.Balance += amount
    return nil
}

func (a *Account) Withdraw(amount int64) error {
    if amount <= 0 {
        return errors.New("withdrawal amount must be positive")
    }
    if amount > a.Balance {
        return fmt.Errorf("insufficient funds: have %d, need %d", a.Balance, amount)
    }
    a.Balance -= amount
    return nil
}

func (a *Account) Transfer(to *Account, amount int64) error {
    if to == nil {
        return errors.New("destination account cannot be nil")
    }
    if a.Currency != to.Currency {
        return errors.New("currency mismatch")
    }
    
    if err := a.Withdraw(amount); err != nil {
        return fmt.Errorf("transfer failed: %w", err)
    }
    
    if err := to.Deposit(amount); err != nil {
        // Rollback the withdrawal
        a.Deposit(amount)
        return fmt.Errorf("transfer failed: %w", err)
    }
    
    return nil
}

// main.go
func main() {
    acc1 := &bank.Account{ID: "001", Balance: 10000, Currency: "USD"}
    acc2 := &bank.Account{ID: "002", Balance: 5000, Currency: "USD"}
    
    fmt.Printf("Initial balances: %d, %d\n", acc1.Balance, acc2.Balance)
    
    err := acc1.Transfer(acc2, 2000)
    if err != nil {
        fmt.Printf("Transfer failed: %v\n", err)
    } else {
        fmt.Printf("After transfer: %d, %d\n", acc1.Balance, acc2.Balance)
    }
}
```

## Session 2: Concurrency Foundations

### Practice Problem Answers

#### Problem 1: Goroutine Basics
```go
func main() {
    go fmt.Println("1")
    fmt.Println("2")
}
```
**Answer:** C) 2
**Explanation:** Main function doesn't wait for goroutines, so it prints "2" and exits before goroutine can print "1"

#### Problem 2: Channel Deadlock
```go
func main() {
    ch := make(chan int)
    ch <- 42
    val := <-ch
    fmt.Println(val)
}
```
**Answer:** Yes, this will deadlock
**Explanation:** Unbuffered channel blocks on send until there's a receiver. Since send happens before receive in same goroutine, it deadlocks.

#### Problem 3: Buffered vs Unbuffered
**Answer:**
- `ch1` (unbuffered): Synchronous, blocks until receiver ready
- `ch2` (buffered): Asynchronous, can hold 1 value without blocking

#### Problem 4: Worker Pool Implementation
```go
func workerPool(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        results <- job * job  // Square the number
    }
}

func processNumbers(numbers []int) []int {
    jobs := make(chan int, len(numbers))
    results := make(chan int, len(numbers))
    
    var wg sync.WaitGroup
    
    // Start 3 workers
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go workerPool(jobs, results, &wg)
    }
    
    // Send jobs
    for _, num := range numbers {
        jobs <- num
    }
    close(jobs)
    
    // Wait for workers to finish
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    var output []int
    for result := range results {
        output = append(output, result)
    }
    
    return output
}
```

#### Problem 5: Race Condition Fix
```go
var (
    counter int64
    mu      sync.Mutex
)

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            atomic.AddInt64(&counter, 1) // Using atomic
            // OR: 
            // mu.Lock()
            // counter++
            // mu.Unlock()
        }()
    }
    
    wg.Wait()
    fmt.Println(atomic.LoadInt64(&counter)) // Always prints 1000
}
```

### Interview Question Answers

**Q: "What's the difference between buffered and unbuffered channels?"**
**A:** Unbuffered channels are synchronous (block until receiver ready), buffered channels are asynchronous until buffer is full.

**Q: "How do you prevent goroutine leaks?"**
**A:** 
- Always ensure goroutines can exit (use context for cancellation)
- Close channels when done sending
- Use buffered channels to avoid blocking
- Don't start goroutines in loops without bounds

**Q: "When would you use a Mutex vs atomic operations?"**
**A:** 
- Mutex: protecting multiple operations, complex critical sections
- Atomic: simple operations on single values, better performance

**Q: "What happens when you close a channel?"**
**A:** 
- No more values can be sent (panic if attempted)
- Receivers get zero value and ok=false
- Range loops will exit
- Multiple closes panic

**Q: "Explain the difference between concurrency and parallelism."**
**A:** 
- Concurrency: dealing with multiple things at once (composition)
- Parallelism: doing multiple things at once (execution)

**Q: "How do you handle timeouts in Go?"**
**A:** Use `context.WithTimeout`, `time.After` in select, or `http.Client.Timeout`

**Q: "What's the purpose of the select statement?"**
**A:** Non-blocking communication on multiple channels, implementing timeouts, graceful shutdown

**Q: "How do you wait for multiple goroutines to complete?"**
**A:** Use `sync.WaitGroup` - Add before starting goroutines, Done in each goroutine, Wait to block until all complete

## Session 3: Package Design & Architecture

### Practice Problem Answers

#### Problem 1: Refactor Layered Architecture
**Before (MVC):**
```
controllers/user_controller.go
models/user.go
services/user_service.go
repositories/user_repository.go
```

**After (Domain-driven):**
```
userservice/
├── user.go          // Domain types
├── service.go       // Business logic
├── store.go         // Data access interface
├── http.go          // HTTP handlers
└── user_test.go     // Tests

internal/
├── postgres/
│   └── user_store.go // Store implementation
└── config/
    └── config.go
```

#### Problem 2: Resolve Cyclic Dependency
**Solution using interface segregation:**
```go
// blog/post.go
package blog

type AuthorInfo interface {
    GetID() string
    GetName() string
}

type Post struct {
    ID     string
    Author AuthorInfo  // Interface, not concrete type
}

// user/profile.go  
package user

type PostSummary interface {
    GetID() string
    GetTitle() string
}

type User struct {
    ID    string
    Posts []PostSummary  // Interface, not concrete type
}

func (u User) GetID() string   { return u.ID }
func (u User) GetName() string { return u.Name }

// In blog package, Post satisfies PostSummary
// In user package, User satisfies AuthorInfo
```

#### Problem 3: Interface Design
```go
// Catalog service
type ProductCatalog interface {
    GetProduct(id string) (*Product, error)
    ListProducts(category string) ([]*Product, error)
}

// Cart service  
type ShoppingCart interface {
    AddItem(productID string, quantity int) error
    RemoveItem(productID string) error
    GetItems() []CartItem
}

// Order processing
type OrderProcessor interface {
    CreateOrder(cart ShoppingCart) (*Order, error)
    GetOrder(id string) (*Order, error)
}

// Payment handling
type PaymentProcessor interface {
    ProcessPayment(amount int64, method PaymentMethod) (*PaymentResult, error)
}

// User management
type UserManager interface {
    GetUser(id string) (*User, error)
    CreateUser(email, name string) (*User, error)
}
```

#### Problem 4: Package Organization
**Refactored structure:**
```go
// ecommerce/user.go
package ecommerce

type User struct {
    ID    string
    Email string
    Name  string
}

type UserService struct {
    store UserStore
}

type UserStore interface {
    GetUser(id string) (*User, error)
    SaveUser(*User) error
}

// ecommerce/product.go  
type Product struct {
    ID    string
    Name  string
    Price int64
}

type ProductService struct {
    store ProductStore
}

// ecommerce/order.go
type Order struct {
    ID       string
    UserID   string
    Products []OrderItem
}

type OrderService struct {
    userSvc    *UserService
    productSvc *ProductService
    paymentSvc PaymentProcessor
}

// And so on for each domain...
```

### Interview Question Answers

**Q: "How do you organize a Go project?"**
**A:** 
- Start simple with single package
- Create domain-driven packages as complexity grows
- Use `internal/` for private packages
- Separate `cmd/` for multiple binaries
- Avoid technical layering (MVC patterns)

**Q: "How do you handle dependencies between packages?"**
**A:**
- Define interfaces in consumer packages
- Use dependency injection
- Accept interfaces, return structs
- Extract shared types if needed
- Use events for decoupling

**Q: "What causes cyclic dependencies and how do you fix them?"**
**A:**
- Caused by mutual dependencies between packages
- Fix with: interface segregation, shared types package, dependency inversion
- Prevention: proper domain boundaries, dependency direction rules

**Q: "Why doesn't Go have a traditional MVC framework?"**
**A:**
- Go prefers composition over inheritance
- Packages as domain boundaries, not technical layers
- Standard library demonstrates this pattern
- Better testability and maintainability
- Explicit dependencies over magic frameworks

**Q: "What's the difference between `internal/` and `pkg/` directories?"**
**A:**
- `internal/`: Private packages, cannot be imported by external modules
- `pkg/`: Public packages that can be imported by other projects
- `internal/` enforced by Go compiler, `pkg/` is convention

**Q: "How do you design interfaces in Go?"**
**A:**
- Small, focused interfaces (1-3 methods)
- Define where they're used, not where they're implemented
- Accept interfaces, return structs
- Compose larger interfaces from smaller ones

**Q: "What's the best way to structure tests in Go?"**
**A:**
- Unit tests alongside source files (*_test.go)
- Integration tests with build tags
- Table-driven tests for multiple scenarios
- External test packages (package_test) for black-box testing
- Mock implementations for dependencies

**Q: "How do you handle mocking in Go tests?"**
**A:**
- Define interfaces for dependencies
- Create manual mocks or use tools like mockgen
- Inject dependencies through interfaces
- Verify method calls and return controlled responses

## Session 4: Standard Library & Interview Simulation

### Modern Go Features Answers

#### Embedded Files Questions
**Q: "How do you embed files in a Go binary?"**
**A:** Use `//go:embed` directive with `embed` package. Can embed as string, []byte, or embed.FS

**Q: "What's the difference between embedding as string vs []byte vs embed.FS?"**
**A:** 
- `string`: Text files, UTF-8 encoded
- `[]byte`: Binary files, raw bytes
- `embed.FS`: Multiple files/directories, provides filesystem interface

**Q: "When would you use go:generate?"**
**A:** For code generation tasks: enum string methods (stringer), mocks (mockgen), database code (sqlc), protobuf compilation

**Q: "How does stringer work and when is it useful?"**
**A:** Generates String() methods for integer-based types. Useful for enums, status codes, error types to improve debugging and logging

**Q: "What are the benefits of using sqlc over hand-written database code?"**
**A:** Type safety, compile-time SQL validation, reduced boilerplate, automatic struct generation, no reflection overhead

**Q: "How do you handle embedded files in tests?"**
**A:** Same as production code, but can also use `embed.FS` with `testing/fstest` for more complex scenarios

### Problem Solutions

#### Problem 1: Embedded Files Web Server
```go
package main

import (
    "embed"
    "encoding/json"
    "net/http"
    "log"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed templates/*
var templates embed.FS

//go:embed config.json
var configData []byte

type Config struct {
    AppName string `json:"app_name"`
    Version string `json:"version"`
    Debug   bool   `json:"debug"`
}

func main() {
    // Serve static files
    http.Handle("/static/", http.FileServer(http.FS(staticFiles)))
    
    // Serve templates
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := templates.ReadFile("templates/index.html")
        if err != nil {
            http.Error(w, "Template not found", 404)
            return
        }
        w.Header().Set("Content-Type", "text/html")
        w.Write(tmpl)
    })
    
    // Health check with embedded config
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        var config Config
        if err := json.Unmarshal(configData, &config); err != nil {
            http.Error(w, "Config error", 500)
            return
        }
        
        response := map[string]interface{}{
            "status": "healthy",
            "config": config,
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    })
    
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

#### Problem 2: HTTP Status Categories with Stringer
```go
package main

//go:generate stringer -type=StatusCategory

import "fmt"

type StatusCategory int

const (
    Informational StatusCategory = iota + 1 // 1xx
    Success                                 // 2xx
    Redirection                            // 3xx
    ClientError                            // 4xx
    ServerError                            // 5xx
)

func main() {
    status := ClientError
    fmt.Printf("Status category: %s\n", status.String()) // "ClientError"
}
```

Generated `statuscategory_string.go`:
```go
// Code generated by "stringer -type=StatusCategory"; DO NOT EDIT.

package main

import "strconv"

func _() {
    var x [1]struct{}
    _ = x[Informational-1]
    _ = x[Success-2]
    _ = x[Redirection-3]
    _ = x[ClientError-4]
    _ = x[ServerError-5]
}

const _StatusCategory_name = "InformationalSuccessRedirectionClientErrorServerError"

var _StatusCategory_index = [...]uint8{0, 13, 20, 31, 42, 53}

func (i StatusCategory) String() string {
    i -= 1
    if i < 0 || i >= StatusCategory(len(_StatusCategory_index)-1) {
        return "StatusCategory(" + strconv.Itoa(int(i+1)) + ")"
    }
    return _StatusCategory_name[_StatusCategory_index[i]:_StatusCategory_index[i+1]]
}
```

#### Problem 3: Blog Database with sqlc
**schema.sql:**
```sql
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

**queries.sql:**
```sql
-- name: CreatePost :one
INSERT INTO posts (title, content, author_id)
VALUES ($1, $2, $3)
RETURNING id, title, content, author_id, created_at;

-- name: GetPost :one
SELECT id, title, content, author_id, created_at
FROM posts
WHERE id = $1;

-- name: ListPosts :many
SELECT id, title, content, author_id, created_at
FROM posts
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET title = $2, content = $3
WHERE id = $1
RETURNING id, title, content, author_id, created_at;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;
```

**Service layer:**
```go
package blog

import (
    "context"
    "database/sql"
    "myapp/internal/db"
)

type Service struct {
    queries *db.Queries
}

func NewService(database *sql.DB) *Service {
    return &Service{
        queries: db.New(database),
    }
}

func (s *Service) CreatePost(ctx context.Context, title, content string, authorID int32) (*db.Post, error) {
    return s.queries.CreatePost(ctx, db.CreatePostParams{
        Title:    title,
        Content:  content,
        AuthorID: authorID,
    })
}

func (s *Service) GetPost(ctx context.Context, id int32) (*db.Post, error) {
    post, err := s.queries.GetPost(ctx, id)
    if err != nil {
        return nil, err
    }
    return &post, nil
}

func (s *Service) ListPosts(ctx context.Context, limit, offset int32) ([]db.Post, error) {
    return s.queries.ListPosts(ctx, db.ListPostsParams{
        Limit:  limit,
        Offset: offset,
    })
}
```

#### Problem 4: Rate Limiter Solution
```go
package main

import (
    "sync"
    "time"
)

type TokenBucket struct {
    capacity    int64
    tokens      int64
    refillRate  int64
    lastRefill  time.Time
    mu          sync.Mutex
}

func NewTokenBucket(capacity, refillRate int64) *TokenBucket {
    return &TokenBucket{
        capacity:   capacity,
        tokens:     capacity,
        refillRate: refillRate,
        lastRefill: time.Now(),
    }
}

func (tb *TokenBucket) Allow() bool {
    tb.mu.Lock()
    defer tb.mu.Unlock()
    
    now := time.Now()
    elapsed := now.Sub(tb.lastRefill)
    tokensToAdd := int64(elapsed.Seconds()) * tb.refillRate
    
    tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)
    tb.lastRefill = now
    
    if tb.tokens > 0 {
        tb.tokens--
        return true
    }
    
    return false
}

type RateLimiter struct {
    buckets map[string]*TokenBucket
    capacity int64
    refillRate int64
    mu      sync.RWMutex
}

func NewRateLimiter(capacity, refillRate int64) *RateLimiter {
    return &RateLimiter{
        buckets:    make(map[string]*TokenBucket),
        capacity:   capacity,
        refillRate: refillRate,
    }
}

func (rl *RateLimiter) Allow(clientID string) bool {
    rl.mu.RLock()
    bucket, exists := rl.buckets[clientID]
    rl.mu.RUnlock()
    
    if !exists {
        rl.mu.Lock()
        // Double-check after acquiring write lock
        if bucket, exists = rl.buckets[clientID]; !exists {
            bucket = NewTokenBucket(rl.capacity, rl.refillRate)
            rl.buckets[clientID] = bucket
        }
        rl.mu.Unlock()
    }
    
    return bucket.Allow()
}

func min(a, b int64) int64 {
    if a < b {
        return a
    }
    return b
}
```

#### Problem 5: Generic Cache with TTL
```go
package main

import (
    "sync"
    "time"
)

type cacheItem[V any] struct {
    value  V
    expiry time.Time
}

type Cache[K comparable, V any] struct {
    items map[K]*cacheItem[V]
    mu    sync.RWMutex
}

func NewCache[K comparable, V any]() *Cache[K, V] {
    c := &Cache[K, V]{
        items: make(map[K]*cacheItem[V]),
    }
    
    // Start cleanup goroutine
    go c.cleanup()
    
    return c
}

func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    expiry := time.Now().Add(ttl)
    c.items[key] = &cacheItem[V]{
        value:  value,
        expiry: expiry,
    }
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    item, exists := c.items[key]
    if !exists {
        var zero V
        return zero, false
    }
    
    if time.Now().After(item.expiry) {
        var zero V
        return zero, false
    }
    
    return item.value, true
}

func (c *Cache[K, V]) Delete(key K) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    delete(c.items, key)
}

func (c *Cache[K, V]) cleanup() {
    ticker := time.NewTicker(time.Minute)
    defer ticker.Stop()
    
    for range ticker.C {
        c.mu.Lock()
        now := time.Now()
        for key, item := range c.items {
            if now.After(item.expiry) {
                delete(c.items, key)
            }
        }
        c.mu.Unlock()
    }
}

// Usage example
func main() {
    cache := NewCache[string, int]()
    
    cache.Set("key1", 42, 5*time.Second)
    
    if value, found := cache.Get("key1"); found {
        fmt.Println("Found:", value)
    }
    
    time.Sleep(6 * time.Second)
    
    if _, found := cache.Get("key1"); !found {
        fmt.Println("Key expired")
    }
}
```

### Mock Interview Code Review Solutions

#### Code Review Problem 1
```go
// What's wrong with this code?
func ProcessUsers(users []User) error {
    for _, user := range users {
        go func() {
            if err := processUser(user); err != nil {
                return err // Wrong!
            }
        }()
    }
    return nil
}
```

**Issues:**
1. **Variable capture**: All goroutines capture the same `user` variable (last iteration value)
2. **Error handling**: `return err` inside goroutine doesn't return from `ProcessUsers`
3. **No synchronization**: Function returns immediately without waiting for goroutines
4. **Lost errors**: No way to collect errors from goroutines

**Fixed version:**
```go
func ProcessUsers(users []User) error {
    var wg sync.WaitGroup
    errCh := make(chan error, len(users))
    
    for _, user := range users {
        wg.Add(1)
        go func(u User) { // Pass user as parameter
            defer wg.Done()
            if err := processUser(u); err != nil {
                errCh <- err // Send error to channel
            }
        }(user)
    }
    
    // Wait for all goroutines and close error channel
    go func() {
        wg.Wait()
        close(errCh)
    }()
    
    // Collect all errors
    var errors []error
    for err := range errCh {
        errors = append(errors, err)
    }
    
    if len(errors) > 0 {
        return fmt.Errorf("processing failed: %v", errors)
    }
    
    return nil
}
```

### System Design Question Examples

#### Q: "Design a distributed task queue system in Go"

**Key Components:**
```go
// Task definition
type Task struct {
    ID       string    `json:"id"`
    Type     string    `json:"type"`
    Payload  []byte    `json:"payload"`
    Priority int       `json:"priority"`
    Retry    int       `json:"retry"`
    Deadline time.Time `json:"deadline"`
}

// Queue interface
type Queue interface {
    Enqueue(task *Task) error
    Dequeue() (*Task, error)
    Ack(taskID string) error
    Nack(taskID string) error
}

// Worker interface  
type Worker interface {
    Process(task *Task) error
}

// Broker coordinates everything
type Broker struct {
    queue   Queue
    workers []Worker
    // ... other components
}
```

**Answer should cover:**
- Queue persistence (Redis/database)
- Worker pool management
- Retry mechanisms with exponential backoff
- Dead letter queues
- Health checks and monitoring
- Horizontal scaling strategies

#### Q: "How would you implement a circuit breaker pattern?"

```go
type CircuitBreaker struct {
    maxFailures  int
    resetTimeout time.Duration
    state        State
    failures     int
    lastFailTime time.Time
    mu           sync.RWMutex
}

type State int

const (
    Closed State = iota
    Open
    HalfOpen
)

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    if cb.state == Open {
        if time.Since(cb.lastFailTime) > cb.resetTimeout {
            cb.state = HalfOpen
            cb.failures = 0
        } else {
            return errors.New("circuit breaker is open")
        }
    }
    
    err := fn()
    
    if err != nil {
        cb.failures++
        cb.lastFailTime = time.Now()
        
        if cb.failures >= cb.maxFailures {
            cb.state = Open
        }
        return err
    }
    
    // Success - reset if we were half-open
    if cb.state == HalfOpen {
        cb.state = Closed
    }
    cb.failures = 0
    
    return nil
}
```

### Final Interview Preparation Summary

#### Essential Go Knowledge Checklist ✅

**Core Language:**
- [ ] Variables, types, and zero values
- [ ] Pointers and memory model
- [ ] Arrays, slices, maps, structs
- [ ] Control flow (if, switch, for, range)
- [ ] Functions (multiple returns, closures, defer)
- [ ] Interfaces and type system
- [ ] Error handling patterns

**Concurrency:**
- [ ] Goroutines and channels
- [ ] Select statements and timeouts
- [ ] sync package (Mutex, WaitGroup, atomic)
- [ ] Context for cancellation
- [ ] Common concurrency patterns

**Package Design:**
- [ ] Domain-driven package organization
- [ ] Dependency inversion and interface design
- [ ] Avoiding cyclic dependencies
- [ ] Testing strategies (unit, integration, mocks)

**Standard Library & Modern Features:**
- [ ] HTTP client/server patterns
- [ ] JSON marshaling/unmarshaling
- [ ] File I/O and string operations
- [ ] Embedded files (go:embed)
- [ ] Code generation (go:generate)
- [ ] Basic generics usage

**Production Patterns:**
- [ ] Error wrapping and unwrapping
- [ ] Graceful shutdown
- [ ] Configuration management
- [ ] Logging and monitoring
- [ ] Performance considerations

#### Last-Minute Interview Tips

1. **Start Simple, Then Optimize**: Begin with working solution, then improve
2. **Think Out Loud**: Explain your reasoning as you code
3. **Ask Clarifying Questions**: Understand requirements before coding
4. **Handle Edge Cases**: Consider nil values, empty inputs, error conditions
5. **Test Your Code**: Walk through examples mentally
6. **Know When to Stop**: Don't over-engineer unless asked

#### Common Mistakes to Avoid

1. **Ignoring errors** - Always handle them properly
2. **Race conditions** - Be careful with shared state
3. **Goroutine leaks** - Ensure goroutines can exit
4. **Wrong receiver types** - Use pointers when modifying
5. **Premature optimization** - Focus on correctness first
6. **Not using interfaces** - Accept interfaces, return structs
7. **Poor package design** - Think domains, not layers

---

**You're now fully prepared for your Go interview! Remember: Go values simplicity, clarity, and explicit error handling. When in doubt, choose the approach that's easiest to understand and maintain.**
        