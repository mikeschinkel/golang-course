# Specific Homework Assignments

Complete these tasks before Session 2. **Create a GitHub repository** and push your solutions.

## Assignment 1: Data Types Mastery (30 minutes)
Create `homework1.go` that demonstrates:
1. Declare variables using all 4 declaration methods
2. Create a constant block using `iota` for file permissions (Read=4, Write=2, Execute=1)
3. Demonstrate zero values for 5 different types
4. Show type conversion between `int`, `float64`, and `string`

**Deliverable:** Working Go file with comments explaining each concept.

## Assignment 2: Collection Operations (45 minutes)
Create `homework2.go` with these functions:
1. `reverseSlice([]int) []int` - reverse a slice without using built-in functions
2. `mergeMaps(map[string]int, map[string]int) map[string]int` - merge two maps, second overwrites first
3. `findDuplicates([]string) []string` - return duplicate elements
4. `groupByLength([]string) map[int][]string` - group strings by length

**Deliverable:** All functions with test cases showing they work correctly.

## Assignment 3: Pointer Practice (30 minutes)
Create `homework3.go` implementing:
1. A `swap(a, b *int)` function that swaps two integers
2. A `Node` struct for a linked list with `Next *Node` field
3. Functions to add/remove nodes from the linked list
4. Demonstrate the difference between value and pointer receivers

**Deliverable:** Working linked list with at least 3 operations.

## Assignment 4: Error Handling (20 minutes)
Create `homework4.go` with:
1. A custom error type `ValidationError` with field name and message
2. A function `validateUser(name, email string) error` that returns your custom error
3. A function that wraps errors using `fmt.Errorf` with `%w`
4. Demonstrate error unwrapping with `errors.Is` and `errors.As`

**Deliverable:** Complete error handling examples with different error types.

## Assignment 5: Mini Project - Bank Account (60 minutes)
Create `bank/` package with:
1. `Account` struct with fields: ID, Balance, Currency
2. Methods: `Deposit(amount)`, `Withdraw(amount)`, `Transfer(to *Account, amount)`
3. All methods must handle errors appropriately
4. Use pointer receivers where needed
5. Create a simple `main.go` that demonstrates all operations

**Deliverable:** Complete package with main function showing all operations.

## Verification Checklist
Before Session 2, ensure you can:
- [ ] Explain when to use pointer vs value receivers
- [ ] Create and manipulate slices, maps, and structs
- [ ] Handle errors idiomatically
- [ ] Use all Go control structures
- [ ] Implement interfaces implicitly
- [ ] Debug common pointer mistakes

**Submit:** GitHub repository URL with all homework files organized in folders.

---

**Session 2 Preview:** We'll review your homework solutions and dive deep into goroutines, channels, and concurrent programming patterns.