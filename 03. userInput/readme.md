# User Input Handling in Go

## Table of Contents
- [Core Concepts](#core-concepts)
- [Input Methods](#input-methods)
- [Data Type Handling](#data-type-handling)
- [Input Validation](#input-validation)
- [Advanced Features](#advanced-features)
- [Best Practices](#best-practices)
- [Common Use Cases](#common-use-cases)
- [Error Handling](#error-handling)
- [Performance Considerations](#performance-considerations)

## Core Concepts

### bufio.NewReader
The `bufio.NewReader` is the foundation of input handling in Go. It provides buffered input operations which are more efficient than reading directly from os.Stdin.

```go
reader := bufio.NewReader(os.Stdin)
```

Key advantages:
- Buffered reading reduces system calls
- Efficient handling of large input streams
- Support for various delimiter-based reading

### Input Stream Management
Go manages input streams through various methods:
- `ReadString('\n')`: Reads until newline
- `ReadBytes('\n')`: Similar to ReadString but returns bytes
- `ReadLine()`: Low-level line reading
- `Scanner`: Alternative for line-by-line reading

## Input Methods

### 1. String Input
```go
input, err := reader.ReadString('\n')
input = strings.TrimSpace(input)  // Remove newline and whitespace
```

Use cases:
- User names
- Text input
- Commands
- File paths

### 2. Numeric Input
```go
numStr, _ := reader.ReadString('\n')
num, err := strconv.ParseFloat(strings.TrimSpace(numStr), 64)
```

Considerations:
- Always handle parsing errors
- Validate range if applicable
- Consider using specific numeric types (int, float32, float64)

### 3. Multiple Value Input
```go
fmt.Print("Enter values (space-separated): ")
input, _ := reader.ReadString('\n')
values := strings.Fields(strings.TrimSpace(input))
```

## Data Type Handling

### String Processing
Common string operations for input:
- `strings.TrimSpace()`: Remove whitespace and newlines
- `strings.Split()`: Break into substrings
- `strings.Fields()`: Split on whitespace
- `strings.ToLower()/ToUpper()`: Case conversion

### Numeric Conversion
Available conversion functions:
```go
strconv.Atoi()           // String to int
strconv.ParseInt()       // String to int64 with base
strconv.ParseFloat()     // String to float64
strconv.ParseBool()      // String to bool
```

## Input Validation

### 1. Basic Validation
```go
if len(input) == 0 {
    return errors.New("input cannot be empty")
}
```

### 2. Type-Specific Validation
```go
// Numeric Range Validation
if num < minValue || num > maxValue {
    return fmt.Errorf("value must be between %v and %v", minValue, maxValue)
}

// String Pattern Validation
if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(input) {
    return errors.New("input must be alphanumeric")
}
```

### 3. Complex Validation
```go
type Validator interface {
    Validate(string) error
}

type InputValidator struct {
    MinLength int
    MaxLength int
    Pattern   string
}
```

## Advanced Features

### 1. ANSI Color Output
```go
const (
    ColorRed    = "\033[31m"
    ColorGreen  = "\033[32m"
    ColorReset  = "\033[0m"
)

fmt.Printf("%sError: %s%s\n", ColorRed, message, ColorReset)
```

### 2. Progress Indication
```go
func showProgress(current, total int) {
    percentage := float64(current) / float64(total) * 100
    fmt.Printf("\rProgress: %.2f%%", percentage)
}
```

### 3. Password Input
```go
func checkPasswordStrength(password string) bool {
    hasNumber := false
    hasUpper := false
    hasLower := false
    hasSpecial := false
    
    for _, char := range password {
        switch {
        case unicode.IsNumber(char):
            hasNumber = true
        case unicode.IsUpper(char):
            hasUpper = true
        case unicode.IsLower(char):
            hasLower = true
        case unicode.IsPunct(char) || unicode.IsSymbol(char):
            hasSpecial = true
        }
    }
    
    return hasNumber && hasUpper && hasLower && hasSpecial
}
```

## Best Practices

### 1. Error Handling
- Always check for errors when reading input
- Provide clear error messages
- Consider retry mechanisms for invalid input
- Log errors appropriately

### 2. Input Buffering
- Use appropriate buffer sizes
- Clear buffers when necessary
- Handle overflow situations

### 3. User Experience
- Provide clear prompts
- Implement timeout mechanisms when appropriate
- Give feedback for long operations
- Handle interrupt signals

### 4. Security Considerations
- Sanitize input data
- Implement timeouts for sensitive operations
- Validate input length and content
- Handle sensitive data appropriately

## Common Use Cases

### 1. Command Line Interface
```go
for {
    fmt.Print("> ")
    cmd, _ := reader.ReadString('\n')
    cmd = strings.TrimSpace(cmd)
    
    switch cmd {
    case "exit":
        return
    case "help":
        showHelp()
    default:
        processCommand(cmd)
    }
}
```

### 2. Form Input
```go
type UserForm struct {
    Name     string
    Age      int
    Email    string
    Password string
}

func collectFormInput() UserForm {
    // Implementation
}
```

### 3. Configuration Input
```go
type Config struct {
    ServerPort int
    DbURL      string
    LogLevel   string
}

func readConfig() Config {
    // Implementation
}
```

## Performance Considerations

### 1. Buffer Sizing
- Default buffer size is usually sufficient
- Increase for large input streams
- Consider memory constraints

### 2. Input Processing
- Process input in chunks for large datasets
- Use goroutines for parallel processing
- Implement appropriate timeouts

### 3. Memory Management
- Clear large buffers after use
- Use appropriate data structures
- Consider string builder for large string operations

## Advanced Implementation Notes

### 1. Concurrent Input Handling
```go
func handleConcurrentInput(inputs chan string, done chan bool) {
    for input := range inputs {
        // Process input
        if shouldStop(input) {
            done <- true
            return
        }
    }
}
```

### 2. Custom Input Readers
```go
type CustomReader struct {
    *bufio.Reader
    timeout time.Duration
}

func (r *CustomReader) ReadWithTimeout() (string, error) {
    ch := make(chan string)
    go func() {
        input, _ := r.ReadString('\n')
        ch <- input
    }()
    
    select {
    case input := <-ch:
        return input, nil
    case <-time.After(r.timeout):
        return "", errors.New("input timeout")
    }
}
```

This implementation provides a robust foundation for handling user input in Go applications, with considerations for performance, security, and user experience. The code can be extended or modified based on specific requirements while maintaining these core principles.