# Type Conversions in Go

## Table of Contents
- [Basic Type Conversions](#basic-type-conversions)
- [String Conversions](#string-conversions)
- [Numeric Conversions](#numeric-conversions)
- [Boolean Conversions](#boolean-conversions)
- [Advanced Conversions](#advanced-conversions)
- [Best Practices](#best-practices)
- [Common Pitfalls](#common-pitfalls)
- [Performance Considerations](#performance-considerations)

## Basic Type Conversions

### Type Casting
Basic type casting in Go uses the type conversion syntax:

```go
value := TypeName(expression)
```

Common type casts:
```go
intVal := int(floatVal)    // float64 to int
float64Val := float64(intVal)  // int to float64
```

When to use:
- Simple numeric type conversions
- Known value ranges
- No risk of data loss

### Type Assertion
Used with interfaces:
```go
value, ok := interface{}.(Type)
```

Use cases:
- Working with interfaces
- Runtime type checking
- Safe type conversion

## String Conversions

### String to Number

1. String to Integer:
```go
// Using Atoi (string to int)
num, err := strconv.Atoi("123")

// Using ParseInt (with base and bit size)
num64, err := strconv.ParseInt("123", 10, 64)
```

When to use Atoi:
- Simple integer conversions
- Base-10 numbers only
- Standard int size

When to use ParseInt:
- Custom bases (2, 8, 16, etc.)
- Specific bit sizes
- More control over conversion

2. String to Float:
```go
// Basic float conversion
float64Val, err := strconv.ParseFloat("3.14", 64)

// With precision control
float32Val, err := strconv.ParseFloat("3.14", 32)
```

Best practices:
- Always check for errors
- Consider precision requirements
- Handle decimal point localization

### Number to String

1. Integer to String:
```go
// Using Itoa (int to string)
str := strconv.Itoa(123)

// Using FormatInt (with base)
binStr := strconv.FormatInt(123, 2)  // Binary
hexStr := strconv.FormatInt(123, 16) // Hexadecimal
```

2. Float to String:
```go
// Using FormatFloat
str := strconv.FormatFloat(3.14, 'f', 2, 64)

// Using Sprintf (more flexible)
str := fmt.Sprintf("%.2f", 3.14)
```

Format specifiers:
- 'f': Fixed-point
- 'e': Scientific notation
- 'g': Auto choice between 'e' and 'f'

## Numeric Conversions

### Integer Type Conversions
```go
var (
    i int = 42
    i8 int8 = int8(i)    // Explicit conversion required
    i16 int16 = int16(i)
    i32 int32 = int32(i)
    i64 int64 = int64(i)
)
```

Considerations:
- Check for overflow
- Consider sign (signed vs unsigned)
- Performance implications

### Float Conversions
```go
var (
    f32 float32 = 3.14
    f64 float64 = float64(f32)
)
```

Precision considerations:
- float32: ~6-9 decimal digits
- float64: ~15-17 decimal digits

## Boolean Conversions

### String to Boolean
```go
// Accepted values: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False
boolVal, err := strconv.ParseBool("true")
```

### Boolean to String
```go
str := strconv.FormatBool(true)  // Returns "true" or "false"
```

## Advanced Conversions

### Base Conversions
```go
// Decimal to other bases
binary := strconv.FormatInt(42, 2)   // Base 2
octal := strconv.FormatInt(42, 8)    // Base 8
hex := strconv.FormatInt(42, 16)     // Base 16

// Other bases to decimal
decimal, err := strconv.ParseInt("2A", 16, 64)  // Hex to decimal
```

Common bases:
- 2: Binary
- 8: Octal
- 10: Decimal
- 16: Hexadecimal
- 36: Maximum base

### Time Conversions
```go
// String to Time
timeVal, err := time.Parse("2006-01-02", "2024-02-06")

// Time to string
formatted := time.Now().Format("2006-01-02 15:04:05")

// Unix timestamp
timestamp := time.Now().Unix()
```

Time format reference:
```
Year: 2006
Month: 01
Day: 02
Hour: 15 (24h) or 03 (12h)
Minute: 04
Second: 05
```

### Byte Conversions
```go
// String to bytes
bytes := []byte("Hello")

// Bytes to string
str := string(bytes)
```

Use cases:
- File I/O
- Network communication
- Data encoding

## Best Practices

### 1. Error Handling
Always check for errors in conversions:
```go
if val, err := strconv.Atoi(str); err != nil {
    // Handle error
} else {
    // Use val
}
```

### 2. Range Validation
```go
func isInRange(val, min, max int64) bool {
    return val >= min && val <= max
}

// Usage
if val, err := strconv.ParseInt(str, 10, 64); err == nil {
    if !isInRange(val, math.MinInt32, math.MaxInt32) {
        // Handle out of range
    }
}
```

### 3. Performance Optimization
```go
// Reuse string builder
var sb strings.Builder
sb.Grow(estimatedSize)  // Pre-allocate if size known

// Use appropriate numeric types
var sum int64  // Instead of float64 for whole numbers
```

## Common Pitfalls

### 1. Loss of Precision
```go
float64Val := 123.456
intVal := int(float64Val)  // Truncates to 123
```

### 2. Integer Overflow
```go
var i8 int8 = 127
i8++  // Overflows to -128
```

### 3. Float Comparison
```go
// Don't do this
if f1 == f2 { ... }

// Do this
const epsilon = 1e-9
if math.Abs(f1 - f2) < epsilon { ... }
```

## Performance Considerations

### 1. String Conversion Cost
- String operations are expensive
- Use appropriate buffer sizes
- Consider using sync.Pool for frequent conversions

### 2. Number Type Selection
```go
// Good: Use int for array indices
for i := 0; i < len(slice); i++ {}

// Bad: Unnecessary float
for f := 0.0; f < float64(len(slice)); f++ {}
```

### 3. Memory Allocation
```go
// Avoid repeated conversions
s := strconv.Itoa(n)  // Do once, reuse result

// Preallocate when possible
bytes := make([]byte, 0, expectedSize)
```

### 4. Batch Processing
For large datasets:
```go
// Process in batches
const batchSize = 1000
for i := 0; i < len(data); i += batchSize {
    end := min(i+batchSize, len(data))
    batch := data[i:end]
    // Process batch
}
```

This documentation provides a comprehensive overview of type conversions in Go, focusing on practical implementations and best practices. The examples and guidelines can be adapted based on specific requirements while maintaining performance and reliability.