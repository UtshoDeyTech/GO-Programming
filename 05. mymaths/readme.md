# Mathematics in Go

## Table of Contents
- [Basic Mathematics](#basic-mathematics)
- [Advanced Math Operations](#advanced-math-operations)
- [Trigonometry](#trigonometry)
- [Random Numbers](#random-numbers)
- [Big Numbers](#big-numbers)
- [Mathematical Constants](#mathematical-constants)
- [Geometric Calculations](#geometric-calculations)
- [Best Practices](#best-practices)
- [Performance Considerations](#performance-considerations)

## Basic Mathematics

### Arithmetic Operations
Basic arithmetic in Go:
```go
x, y := 10, 3

addition := x + y      // Addition
subtraction := x - y   // Subtraction
multiplication := x * y // Multiplication
division := x / y      // Integer division
modulus := x % y       // Modulus
```

Key points:
- Integer division truncates decimals
- Use type conversion for float division
- Modulus works only with integers

### Type-Specific Operations
```go
// Integer operations
var a, b int = 15, 4
intDiv := a / b    // Results in 3

// Float operations
var x, y float64 = 15.0, 4.0
floatDiv := x / y  // Results in 3.75
```

## Advanced Math Operations

### Power and Roots
```go
import "math"

// Powers
square := math.Pow(4, 2)    // 4²
cube := math.Pow(4, 3)      // 4³
arbitrary := math.Pow(2, 8) // 2⁸

// Roots
squareRoot := math.Sqrt(16)  // √16
cubeRoot := math.Cbrt(27)    // ∛27
```

Use cases:
- Geometric calculations
- Scientific computations
- Data scaling

### Logarithms and Exponentials
```go
naturalLog := math.Log(10)    // ln(10)
log10 := math.Log10(100)      // log₁₀(100)
log2 := math.Log2(8)          // log₂(8)
exp := math.Exp(2)            // e²
```

Applications:
- Growth calculations
- Signal processing
- Data normalization

## Trigonometry

### Basic Trigonometric Functions
```go
angle := math.Pi / 4  // 45 degrees in radians

sin := math.Sin(angle)
cos := math.Cos(angle)
tan := math.Tan(angle)
```

### Inverse Trigonometric Functions
```go
asin := math.Asin(0.5)  // sin⁻¹(0.5)
acos := math.Acos(0.5)  // cos⁻¹(0.5)
atan := math.Atan(1.0)  // tan⁻¹(1.0)
```

Angle conversions:
```go
// Degrees to radians
radians := degrees * math.Pi / 180.0

// Radians to degrees
degrees := radians * 180.0 / math.Pi
```

## Random Numbers

### Cryptographic Random Numbers
Using crypto/rand for secure random numbers:
```go
import (
    "crypto/rand"
    "math/big"
)

// Generate random number between 0 and max-1
max := big.NewInt(100)
n, err := rand.Int(rand.Reader, max)
```

When to use:
- Security-critical applications
- Cryptographic operations
- Token generation

## Big Numbers

### Big Integers
Using math/big for arbitrary-precision arithmetic:
```go
import "math/big"

// Creating big integers
a := new(big.Int)
a.SetString("123456789123456789", 10)

// Operations
sum := new(big.Int).Add(a, b)
product := new(big.Int).Mul(a, b)
quotient := new(big.Int).Div(a, b)
```

### Big Floats
High-precision floating-point calculations:
```go
// Set precision and value
pi := new(big.Float).SetPrec(200)
pi.SetString("3.14159265358979323846264338327950288419716939937510")

// Operations
result := new(big.Float).Mul(pi, pi)
```

Use cases:
- Financial calculations
- Scientific computing
- Cryptography

## Mathematical Constants

### Built-in Constants
```go
math.Pi      // π (3.14159...)
math.E       // e (2.71828...)
math.Phi     // φ (1.61803...) - Golden ratio
math.Sqrt2   // √2 (1.41421...)
math.Ln2     // ln(2)
math.Log2E   // log₂(e)
math.Log10E  // log₁₀(e)
```

## Geometric Calculations

### 2D Geometry
```go
// Circle calculations
radius := 5.0
area := math.Pi * math.Pow(radius, 2)
circumference := 2 * math.Pi * radius

// Triangle calculations
base := 4.0
height := 3.0
triangleArea := (base * height) / 2
```

### 3D Geometry
```go
// Sphere calculations
radius := 5.0
volume := (4.0/3.0) * math.Pi * math.Pow(radius, 3)
surfaceArea := 4 * math.Pi * math.Pow(radius, 2)

// Cube calculations
side := 3.0
cubeVolume := math.Pow(side, 3)
cubeSurfaceArea := 6 * math.Pow(side, 2)
```

## Best Practices

### 1. Precision Handling
```go
// Avoid direct float comparisons
const epsilon = 1e-10
isEqual := math.Abs(a - b) < epsilon

// Use appropriate types
var sum float64    // Instead of float32 for better precision
```

### 2. Error Checking
```go
// Check for domain errors
if x < 0 {
    return 0, fmt.Errorf("negative input: %v", x)
}

// Handle special cases
if math.IsNaN(x) || math.IsInf(x, 0) {
    return 0, fmt.Errorf("invalid input: %v", x)
}
```

### 3. Performance Optimization
```go
// Pre-calculate constants
const twoPi = 2 * math.Pi

// Use math.Pow sparingly for integer powers
square := x * x        // Instead of math.Pow(x, 2)
cube := x * x * x      // Instead of math.Pow(x, 3)
```

## Performance Considerations

### 1. Type Selection
```go
// Use int for indices and counters
for i := 0; i < len(slice); i++ {}

// Use float64 for most floating-point calculations
var result float64
```

### 2. Memory Management
```go
// Reuse big.Int/big.Float objects
result := new(big.Int)
for i := 0; i < iterations; i++ {
    result.Mul(result, multiplier)
}
```

### 3. Optimization Techniques
```go
// Use bit operations when possible
isEven := (n & 1) == 0    // Instead of n % 2 == 0

// Avoid unnecessary conversions
intVal := int(floatVal)   // Only when needed
```

This documentation provides a comprehensive overview of mathematical operations in Go, focusing on practical implementations and best practices. The examples and guidelines can be adapted based on specific requirements while maintaining performance and reliability.