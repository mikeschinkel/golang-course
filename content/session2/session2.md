# Session 2: Concurrency Foundations

**Duration:** 2 hours | **Focus:** Goroutines, Channels, Synchronization

## Overview

This session dives deep into Go's concurrency model, covering goroutines, channels, and synchronization patterns. You'll learn how to write concurrent programs that are both efficient and safe.

## Session Outline

### 1. [Key Concepts to Master](./01-key-concepts-to-master/)
- [Goroutines](./01-key-concepts-to-master/01-goroutines.md)
- [Channels - Go's Concurrency Primitive](./01-key-concepts-to-master/02-channels-gos-concurrency-primitive.md)
- [Select Statement](./01-key-concepts-to-master/03-select-statement.md)
- [Synchronization Patterns](./01-key-concepts-to-master/04-synchronization-patterns.md)

### 2. [Common Patterns](./02-common-patterns/)
- [Worker Pool](./02-common-patterns/01-worker-pool.md)
- [Fan-out, Fan-in](./02-common-patterns/02-fan-out-fan-in.md)
- [Context for Cancellation](./02-common-patterns/03-context-for-cancellation.md)

### 3. [Practice Problems](./03-practice-problems/)
- [Goroutine Basics](./03-practice-problems/01-goroutine-basics.md)
- [Channel Deadlock](./03-practice-problems/02-channel-deadlock.md)
- [Buffered vs Unbuffered](./03-practice-problems/03-buffered-vs-unbuffered.md)
- [Worker Pool Implementation](./03-practice-problems/04-worker-pool-implementation.md)
- [Race Condition](./03-practice-problems/05-race-condition.md)

### 4. [Common Interview Questions](./04-common-interview-questions.md)

### 5. [Critical Gotchas](./05-critical-gotchas/)
- [Goroutine Variable Capture](./05-critical-gotchas/01-goroutine-variable-capture.md)
- [Channel Direction](./05-critical-gotchas/02-channel-direction.md)
- [Select with Default](./05-critical-gotchas/03-select-with-default.md)

### 6. [Performance Tips](./06-performance-tips.md)

### 7. [Homework for Next Session](./07-homework-for-next-session.md)

## Key Learning Objectives

By the end of this session, you should be able to:

- [ ] Understand the difference between buffered and unbuffered channels
- [ ] Know how to prevent goroutine leaks
- [ ] Understand when to use Mutex vs atomic operations
- [ ] Know what happens when you close a channel
- [ ] Explain the difference between concurrency and parallelism
- [ ] Handle timeouts in Go
- [ ] Understand the purpose of the select statement
- [ ] Know how to wait for multiple goroutines to complete

## Critical Concepts

- **Goroutines** are lightweight threads managed by the Go runtime
- **Channels** provide communication between goroutines
- **Select** enables non-blocking channel operations
- **WaitGroup** helps coordinate goroutine completion
- **Mutex** provides mutual exclusion for shared resources

---

**Previous Session:** [Session 1: Go Fundamentals & Memory Management](../session1/session1.md)  
**Next Session:** [Session 3: Package Design & Go Architecture Patterns](../session3/session3.md)
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
