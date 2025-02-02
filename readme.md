# Complete Go Programming Guide

## Table of Contents
1. [Basic Types](#1-basic-types)
2. [Composite Types](#2-composite-types)
3. [Special Types](#3-special-types)
4. [Core Modules](#4-core-modules)
5. [Functions and Methods](#5-functions-and-methods)
6. [Error Handling](#6-error-handling)
7. [Concurrency](#7-concurrency)
8. [Standard Library](#8-standard-library)
9. [Best Practices](#9-best-practices)

## 1. Basic Types

### 1.1 Numeric Types
#### Integer Types
- Signed Integers:
  - `int`: Platform dependent (32/64 bit)
  - `int8`: -128 to 127
  - `int16`: -32,768 to 32,767
  - `int32`: -2,147,483,648 to 2,147,483,647
  - `int64`: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807

- Unsigned Integers:
  - `uint`: Platform dependent
  - `uint8`: 0 to 255
  - `uint16`: 0 to 65,535
  - `uint32`: 0 to 4,294,967,295
  - `uint64`: 0 to 18,446,744,073,709,551,615
  - `uintptr`: Large enough to store pointer value

#### Floating-Point Types
- `float32`: ±3.4e38 (~6-7 decimal digits precision)
- `float64`: ±1.8e308 (~15-17 decimal digits precision)

#### Complex Types
- `complex64`: Complex numbers with float32 parts
- `complex128`: Complex numbers with float64 parts

### 1.2 Text Types
#### String
- UTF-8 encoded
- Immutable
- Zero value: `""`
- Common operations:
  - Concatenation: `+`
  - Length: `len()`
  - Substring: `string[start:end]`

#### Rune
- Alias for int32
- Represents a Unicode code point
- Used for character manipulation
- Single quotes: `'A'`

### 1.3 Boolean
- Values: `true`, `false`
- Zero value: `false`
- Operators:
  - AND: `&&`
  - OR: `||`
  - NOT: `!`

### 1.4 Byte
- Alias for uint8
- Used for byte-level operations
- Common in I/O operations

## 2. Composite Types

### 2.1 Array
- Fixed-length sequence
- Zero-based indexing
- Declaration: `var name [size]type`
- Examples:
  ```go
  var numbers [5]int
  names := [3]string{"John", "Jane", "Joe"}
  matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}
  ```

### 2.2 Slice
- Dynamic-length sequence
- Built on arrays
- Components:
  - Pointer to underlying array
  - Length
  - Capacity
- Operations:
  - `append()`
  - `copy()`
  - `make()`
  - Slicing: `slice[start:end]`

### 2.3 Map
- Key-value pairs
- Unordered collection
- Declaration: `map[KeyType]ValueType`
- Operations:
  - Insert: `map[key] = value`
  - Delete: `delete(map, key)`
  - Lookup: `value, exists := map[key]`

### 2.4 Struct
- User-defined type
- Groups related data
- Supports embedding
- Example:
  ```go
  type Person struct {
      Name    string
      Age     int
      Address Address
  }
  ```

## 3. Special Types

### 3.1 Pointer
- Stores memory address
- Zero value: `nil`
- Operations:
  - Address of: `&`
  - Dereference: `*`
- Uses:
  - Pass by reference
  - Memory efficiency
  - Mutability

### 3.2 Interface
- Defines method set
- Implicit implementation
- Common interfaces:
  - `io.Reader`
  - `io.Writer`
  - `error`
- Example:
  ```go
  type Writer interface {
      Write([]byte) (int, error)
  }
  ```

### 3.3 Channel
- Goroutine communication
- Types:
  - Unbuffered
  - Buffered
- Operations:
  - Send: `ch <- value`
  - Receive: `value := <-ch`
  - Close: `close(ch)`

## 4. Core Modules

### 4.1 fmt
- Formatting and printing
- Key functions:
  - `Print()`: Basic output
  - `Println()`: Output with newline
  - `Printf()`: Formatted output
  - `Sprintf()`: Format to string
  - `Scanf()`: Formatted input

#### Format Verbs
- General:
  - `%v`: Default format
  - `%T`: Type
  - `%#v`: Go-syntax format
- Integer:
  - `%d`: Decimal
  - `%b`: Binary
  - `%o`: Octal
  - `%x`: Hexadecimal
- Float:
  - `%f`: Decimal point
  - `%e`: Scientific notation
  - `%g`: Compact format
- String:
  - `%s`: String
  - `%q`: Quoted string
- Width and Precision:
  - `%f`: Default
  - `%9f`: Width 9
  - `%.2f`: Precision 2
  - `%9.2f`: Width 9, precision 2

### 4.2 os
- Operating system interface
- File operations:
  - `Open()`
  - `Create()`
  - `Remove()`
  - `ReadFile()`
  - `WriteFile()`
- Environment:
  - `Getenv()`
  - `Setenv()`
  - `Environ()`
- Process:
  - `Exit()`
  - `Args`
  - `Executable()`

### 4.3 io
- Basic I/O interfaces
- Key interfaces:
  - `Reader`
  - `Writer`
  - `Closer`
- Utilities:
  - `Copy()`
  - `CopyN()`
  - `ReadAll()`
  - `WriteString()`

### 4.4 strings
- String manipulation
- Key functions:
  - `Split()`
  - `Join()`
  - `Contains()`
  - `Replace()`
  - `ToUpper()`
  - `ToLower()`
  - `Trim()`
  - `Index()`

### 4.5 time
- Time and duration
- Key types:
  - `Time`
  - `Duration`
  - `Location`
- Functions:
  - `Now()`
  - `Parse()`
  - `Format()`
  - `Add()`
  - `Sub()`

## 5. Functions and Methods

### 5.1 Function Types
- Basic:
  ```go
  func name(params) returnType {
      // body
  }
  ```
- Multiple returns:
  ```go
  func name() (type1, type2) {
      return val1, val2
  }
  ```
- Named returns:
  ```go
  func name() (n1 type1, n2 type2) {
      n1 = val1
      n2 = val2
      return
  }
  ```
- Variadic:
  ```go
  func name(param ...type) {
      // param is a slice
  }
  ```

### 5.2 Methods
- Function with receiver
- Syntax:
  ```go
  func (receiver Type) name(params) returnType {
      // body
  }
  ```
- Pointer receiver:
  ```go
  func (receiver *Type) name(params) returnType {
      // can modify receiver
  }
  ```

### 5.3 Functional Programming
- First-class functions
- Function types
- Closures
- Callbacks
- Examples:
  ```go
  // Function type
  type Handler func(string) error

  // Anonymous function
  func() {
      // body
  }()

  // Closure
  func outer() func() int {
      x := 0
      return func() int {
          x++
          return x
      }
  }
  ```

## 6. Error Handling

### 6.1 Error Interface
```go
type error interface {
    Error() string
}
```

### 6.2 Error Creation
- `errors.New()`
- `fmt.Errorf()`
- Custom errors:
  ```go
  type CustomError struct {
      Code    int
      Message string
  }

  func (e *CustomError) Error() string {
      return fmt.Sprintf("%d: %s", e.Code, e.Message)
  }
  ```

### 6.3 Error Handling Patterns
- Return error:
  ```go
  if err != nil {
      return err
  }
  ```
- Panic/Recover:
  ```go
  defer func() {
      if r := recover(); r != nil {
          // handle panic
      }
  }()
  ```

## 7. Concurrency

### 7.1 Goroutines
- Lightweight threads
- Creation: `go function()`
- Characteristics:
  - Independently executed
  - Share same address space
  - Managed by Go runtime

### 7.2 Channels
- Communication
- Types:
  ```go
  ch := make(chan int)        // unbuffered
  ch := make(chan int, 100)   // buffered
  ```
- Operations:
  ```go
  ch <- x      // send
  x := <-ch    // receive
  close(ch)    // close
  ```

### 7.3 Select
- Multiple channel operations
- Example:
  ```go
  select {
  case x := <-ch1:
      // handle ch1
  case ch2 <- y:
      // handle ch2
  default:
      // no channel ready
  }
  ```

### 7.4 Sync Package
- Mutex
- WaitGroup
- Once
- Cond
- Pool
- Map

## 8. Standard Library

### 8.1 Common Packages
- `encoding/json`
- `net/http`
- `database/sql`
- `crypto`
- `log`
- `regexp`
- `sort`
- `testing`

### 8.2 Collection Operations
- `make()`
- `new()`
- `append()`
- `copy()`
- `delete()`
- `len()`
- `cap()`

### 8.3 Type Conversions
- `strconv.Itoa()`
- `strconv.Atoi()`
- `strconv.ParseFloat()`
- `strconv.ParseInt()`
- Type assertions:
  ```go
  value.(Type)
  ```

## 9. Best Practices

### 9.1 Code Organization
- Package naming
- File structure
- Import organization
- Interface design

### 9.2 Error Handling
- Always check errors
- Error wrapping
- Error types
- Panic appropriately

### 9.3 Performance
- Memory management
- Goroutine usage
- Channel sizing
- Benchmark testing

### 9.4 Testing
- Unit tests
- Table-driven tests
- Benchmarks
- Examples

### 9.5 Documentation
- Package documentation
- Function documentation
- Example code
- GoDoc formatting