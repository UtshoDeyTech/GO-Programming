# Go Variables Guide

This guide provides a comprehensive overview of variable declarations, types, and best practices in Go.

## Table of Contents
- [Variable Declaration Methods](#variable-declaration-methods)
- [Basic Types](#basic-types)
- [Advanced Types](#advanced-types)
- [Constants](#constants)
- [Type Conversion](#type-conversion)
- [Best Practices](#best-practices)

## Variable Declaration Methods

### 1. Explicit Declaration with Type
```go
var name string = "value"
```
**Best for:**
- Package-level variables
- When type clarity is important
- When you want to ensure a specific type
- When declaring but not initializing

### 2. Type Inference (Implicit)
```go
var name = "value"
```
**Best for:**
- When the type is obvious from the value
- Local variables where type is clear from context
- Reducing code verbosity while maintaining readability

### 3. Short Declaration (`:=`)
```go
name := "value"
```
**Best for:**
- Local variables inside functions
- Quick iterations and prototyping
- When you want concise, readable code
- Multiple variable declarations in one line

### 4. Multiple Variable Declaration
```go
var (
    name    string  = "John"
    age     int     = 25
    height  float64 = 1.75
)
```
**Best for:**
- Grouping related variables
- Package-level declarations
- Improving code organization
- When declaring multiple variables of different types

## Basic Types

### Numeric Types
1. **Integers**
   - `int8`, `int16`, `int32`, `int64`: For signed integers
   - `uint8`, `uint16`, `uint32`, `uint64`: For unsigned integers
   - `int`, `uint`: Platform-dependent size (32 or 64 bits)
   
   **Best practices:**
   - Use `int` for general purpose integers
   - Use specific sizes when memory is critical
   - Use unsigned types only when negative values are impossible

2. **Floating Point**
   - `float32`: Single precision
   - `float64`: Double precision
   
   **Best practices:**
   - Default to `float64` for better precision
   - Use `float32` only when memory is critical

3. **Complex Numbers**
   - `complex64`: 32-bit real and imaginary parts
   - `complex128`: 64-bit real and imaginary parts

### String Type
```go
var name string = "value"
```
**Features:**
- UTF-8 encoded
- Immutable
- Can use backticks for raw strings
- Supports concatenation with `+`

### Boolean Type
```go
var isValid bool = true
```

## Advanced Types

### 1. Custom Types
```go
type UserID uint64
type Temperature float64
```
**Best for:**
- Type safety
- Domain-specific logic
- Improving code readability
- Adding methods to basic types

### 2. Struct Types
```go
type Person struct {
    Name string
    Age  int
}
```
**Best for:**
- Grouping related data
- Creating custom data structures
- Object-oriented patterns
- API responses and data models

### 3. Interface Types
```go
type Printable interface {
    ToString() string
}
```
**Best for:**
- Defining behavior contracts
- Enabling polymorphism
- Testing and mocking
- Loose coupling between components

## Constants

### 1. Single Constants
```go
const MaxConnections = 100
```

### 2. Constant Groups
```go
const (
    Pi       = 3.14159
    MaxUsers = 1000
)
```

### 3. Iota Constants
```go
const (
    Sunday = iota
    Monday
    Tuesday
)
```
**Best for:**
- Enumerations
- Sequential values
- Bit flags
- Status codes

## Type Conversion

### 1. Basic Conversion
```go
float64Value := float64(intValue)
```

### 2. String Conversion
```go
strconv.Itoa(intValue)  // int to string
strconv.Atoi(strValue)  // string to int
```

## Best Practices

1. **Variable Naming**
   - Use camelCase for variables
   - Use PascalCase for exported variables
   - Use meaningful, descriptive names
   - Keep names concise but clear

2. **Scope Management**
   - Declare variables close to their use
   - Use the narrowest scope possible
   - Prefer function parameters over package variables
   - Use package variables only when necessary

3. **Type Selection**
   - Use the simplest type that meets requirements
   - Default to `int` for integers
   - Default to `float64` for floating-point
   - Use specific sizes only when needed

4. **Declaration Style**
   - Use `:=` for local variables
   - Use `var` for package-level variables
   - Group related `var` declarations
   - Use constants for fixed values

5. **Memory Considerations**
   - Use smaller types when dealing with large arrays
   - Consider memory alignment in structs
   - Use pointers for large structs in function parameters
   - Be mindful of string concatenation performance

6. **Zero Values**
   - Utilize zero values when appropriate
   - Initialize variables explicitly when zero value is not desired
   - Use pointers when nil is a valid state
   - Document expected zero value behavior

7. **Error Handling**
   - Always check error returns from type conversions
   - Use multiple return values for error handling
   - Don't ignore type conversion errors
   - Provide context in error messages

## Common Pitfalls to Avoid

1. Unused variables (Go compiler will catch these)
2. Shadowing variables in inner scopes
3. Using global variables when not necessary
4. Ignoring type conversion errors
5. Using float types for monetary values
6. Not considering string immutability in loops

## Performance Considerations

1. **String Operations**
   - Use `strings.Builder` for multiple concatenations
   - Preallocate slices when length is known
   - Use raw strings for regular expressions
   - Consider runes for Unicode operations

2. **Numeric Types**
   - Use appropriate sizes for numeric types
   - Consider memory alignment in structs
   - Use integer arithmetic when possible
   - Be aware of floating-point precision limits

This guide serves as a reference for Go variable usage and best practices. Always consider your specific use case when choosing variable types and declaration methods.