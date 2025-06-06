# Session 3: Package Design & Go Architecture Patterns

**Duration:** 2 hours | **Focus:** Idiomatic Go Organization, Domain-Driven Design, Avoiding Anti-Patterns

## Overview

This session focuses on organizing Go code idiomatically, designing packages around business domains rather than technical layers, and implementing comprehensive testing strategies. You'll learn to avoid common architectural anti-patterns and build maintainable Go applications.

## Session Outline

### 1. [Core Philosophy: Packages as Domain Boundaries](./01-core-philosophy/)
- [Anti-Pattern: Technical Layers (MVC)](./01-core-philosophy/01-anti-pattern.md)
- [Idiomatic Go: Specific Domain Packages](./01-core-philosophy/02-idiomatic-go.md)

### 2. [Package Design Principles](./02-package-design-principles/)
- [High Cohesion, Loose Coupling](./02-package-design-principles/01-high-cohesion-loose-coupling.md)
- [Interfaces at Package Boundaries](./02-package-design-principles/02-interfaces-at-package-boundaries.md)
- [Dependency Direction Rules](./02-package-design-principles/03-dependency-direction-rules.md)

### 3. [Resolving Cyclic Dependencies](./03-resolving-cyclic-dependencies/)
- [Problem: Two-Way Dependencies](./03-resolving-cyclic-dependencies/01-problem-two-way-dependencies.md)
- [Solution 1: Extract Common Types](./03-resolving-cyclic-dependencies/02-solution-extract-common-types.md)
- [Solution 2: Interface Segregation](./03-resolving-cyclic-dependencies/03-solution-interface-segregation.md)
- [Solution 3: Dependency Inversion](./03-resolving-cyclic-dependencies/04-solution-dependency-inversion.md)

### 4. [Idiomatic Project Structure](./04-idiomatic-project-structure/)
- [Small Project](./04-idiomatic-project-structure/01-small-project.md)
- [Medium Project](./04-idiomatic-project-structure/02-medium-project.md)
- [Large Project](./04-idiomatic-project-structure/03-large-project.md)

### 5. [Interface Design Patterns](./05-interface-design-patterns/)
- [Accept Interfaces, Return Structs](./05-interface-design-patterns/01-accept-interfaces-return-structs.md)
- [Small, Focused Interfaces](./05-interface-design-patterns/02-small-focused-interfaces.md)
- [Define Interfaces Where They're Used](./05-interface-design-patterns/03-define-interfaces-where-used.md)

### 6. [Comprehensive Testing Strategy](./06-comprehensive-testing-strategy/)
- [Test Organization and Naming](./06-comprehensive-testing-strategy/01-test-organization-and-naming.md)
- [Mock Implementation Patterns](./06-comprehensive-testing-strategy/02-mock-implementation-patterns.md)
- [Integration Tests](./06-comprehensive-testing-strategy/03-integration-tests.md)
- [Benchmark Tests](./06-comprehensive-testing-strategy/04-benchmark-tests.md)
- [Test Utilities and Helpers](./06-comprehensive-testing-strategy/05-test-utilities-and-helpers.md)
- [Testing HTTP Handlers](./06-comprehensive-testing-strategy/06-testing-http-handlers.md)
- [Test Organization Best Practices](./06-comprehensive-testing-strategy/07-test-organization-best-practices.md)

### 7. [Practice Problems](./07-practice-problems/)
- [Refactor Layered Architecture](./07-practice-problems/01-refactor-layered-architecture.md)
- [Resolve Cyclic Dependency](./07-practice-problems/02-resolve-cyclic-dependency.md)
- [Interface Design](./07-practice-problems/03-interface-design.md)
- [Package Organization](./07-practice-problems/04-package-organization.md)

## Assessment & Review

- [Common Interview Questions](./08-common-interview-questions.md)
- [Key Takeaways](./09-key-takeaways.md)
- [Homework for Next Session](./10-homework-for-next-session.md)

## Key Learning Objectives

By the end of this session, you should be able to:

- [ ] Organize Go projects using domain-driven package structure
- [ ] Handle dependencies between packages properly
- [ ] Identify and resolve cyclic dependencies
- [ ] Understand why Go doesn't use traditional MVC frameworks
- [ ] Know the difference between `internal/` and `pkg/` directories
- [ ] Design small, focused interfaces
- [ ] Structure comprehensive test suites
- [ ] Implement proper mocking strategies

## Core Principles

1. **Think domains, not layers** - Package by business capability
2. **Interfaces belong with consumers** - Not with implementers
3. **Small interfaces are better** - Compose when needed
4. **Dependencies point inward** - Core domain has no external deps
5. **Avoid god packages** - Keep packages focused and cohesive
6. **Use internal/ wisely** - Hide implementation details
7. **Learn from stdlib** - Follow established Go patterns
8. **Test comprehensively** - Unit, integration, benchmarks
9. **Mock at boundaries** - Interface-based testing

---

**Previous Session:** [Session 2: Concurrency Foundations](../session2/session2.md)  
**Next Session:** [Session 4: Standard Library Mastery & Interview Simulation](../session4/session4.md)
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
