# goDemo

## Comparison of Go with Other Programming Languages

### Purpose

This document highlights the basic characteristics of Golang compared with both strong type languages (C/C++/Java) and dynamic script languages (Python/JavaScript/Ruby).

---

## Module & Package

### Python

- **Module:** A single `.py` file containing Python definitions and statements.
- **Package:** A directory containing multiple modules and potentially other sub-packages.

### Java

- **Module:** Defined by a `module-info.java` file, declaring its name, dependencies (`requires`), and exported packages (`exports`).
- **Package:** Hierarchical directory structure, each level of the package name corresponds to a subdirectory.

### Go

- **Module:** A collection of Go packages versioned together as a single unit.
  1. Defined by a `go.mod` file at the root of your project.
  2. Tells Go how to resolve imports, track dependencies, and manage versions.
- **Package:** Declared using the `package` keyword. Go package names are concise and do not necessarily match the directory path. Go imports entire packages, not individual types.

**Demo:** [`pkg/struct`](pkg/struct/composition/demo.go)

---

## Inheritance vs Composition

### C++/Java Inheritance

- Subclass inherits/extends base class; class implements interface.

### Duck Typing (Go and Most Dynamic Script Languages)

- "If it walks like a duck and it quacks like a duck, then it must be a duck."

### Go

- No classes, no `extends` or `super`.
- **Struct embedding:** Enables code reuse by embedding one struct into another (composition).
- **Interfaces:** Define behavior by method signatures; allow polymorphism.

| Feature                | Inheritance (Java, C++) | Go's Way (Composition/Interfaces) |
|------------------------|-------------------------|-----------------------------------|
| Inherit fields/methods | Yes                     | Yes (via embedding)               |
| Method overriding      | Yes                     | No (but can shadow)               |
| Polymorphism           | Yes (class hierarchy)   | Yes (interfaces)                  |
| Multiple inheritance   | No (Java) / Yes (C++)   | Yes (multiple embedding)          |

**Demo:** [`pkg/struct`](pkg/struct/composition/demo.go)

---

## Pointer vs Reference

- **Java:** No pointers, only references.
- **C++:** Supports both pointers and references.
- **Go:** Methods can use value or pointer receivers, affecting behavior.

| Use Case                        | Use a...              | Why                                 |
|----------------------------------|-----------------------|--------------------------------------|
| Modify the receiver              | Pointer receiver \*T  | Avoid copying and allow mutation     |
| Read-only copy                   | Value receiver T      | Mutation isn't needed                |
| Large struct                     | Pointer receiver      | Avoid expensive copying              |
| Small and immutable struct       | Value receiver        | Safe and sometimes more efficient    |

**Demo:** [`pkg/pointer`](pkg/pointer/demo/demo.go)

---

## New vs Make

- **C/C++:** `new` & `delete`
- **Java:** `new` creates an object on the heap; garbage collector releases objects (no `delete`)
- **Go:**

| Feature      | `new(T)`                    | `make(T)` (slice, map, channel only) |
|--------------|-----------------------------|--------------------------------------|
| Works with   | All types                   | Only slices, maps, and channels      |
| Returns      | Pointer to zeroed value \*T | Initialized value of type T          |
| Used for     | Allocating memory           | Creating slices, maps, channels      |
| Zero value   | Yes                         | No (returns initialized value)       |

- Go garbage collector frees memory when there are no more references to a value.
- No `delete`, `free`, or `dispose` in Go.

**Demo:** [`pkg/new`](pkg/new/demo.go)

---

## Stack vs Heap

- **C/C++/Java:** `new` creates an object on the heap; reference itself is allocated on the stack.
- **Go:** Escape analysis determines allocation:
  - Used only inside current function → **Stack**
  - Returned or referenced outside → **Heap**

**Demo:** [`pkg/escape_analysis`](pkg/escape_analysis/demo.go)

---

## Visibility

- **C/C++/Java:** `public`/`private`/`protected`
- **Go:**
  - Starts with uppercase (exported): **Public** (any package)
  - Starts with lowercase (unexported): **Private** (same package only)
  - Applies to structs, methods, functions, variables, constants, interfaces.

**Demo:** [`pkg/visibility`](pkg/visibility/main.go)

---

## Generics

- **C++:** Generics via templates.
- **Java vs Go:**

| Feature           | Java Generics                | Go Generics                        |
|-------------------|-----------------------------|------------------------------------|
| Implementation    | Type erasure (compile-time)  | No type erasure (runtime info)     |
| Primitive types   | Not allowed directly         | Allowed                            |
| Constraints       | Bounded type parameters      | Interfaces                         |
| Operator support  | Limited                      | No direct generic operator support |
| Syntax            | `<T>`                        | `[T any]` or `[T comparable]`      |

**Demo:** [`pkg/generics`](pkg/generics/main.go)