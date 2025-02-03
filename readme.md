# Go Programming Course

A comprehensive guide to learning Go programming language, covering fundamental concepts to advanced topics.

## Course Structure

### Basic Concepts
1. **Hello** - Introduction to Go programming and basic syntax

2. **Variables** - Understanding Go's type system including:
   
   #### Numeric Types
   - Signed Integers:
     ```go
     int    // Platform dependent (32/64 bit)
     int8   // -128 to 127
     int16  // -32,768 to 32,767
     int32  // -2,147,483,648 to 2,147,483,647
     int64  // -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
     ```
   - Unsigned Integers:
     ```go
     uint    // Platform dependent
     uint8   // 0 to 255
     uint16  // 0 to 65,535
     uint32  // 0 to 4,294,967,295
     uint64  // 0 to 18,446,744,073,709,551,615
     ```
   - Floating-Point:
     ```go
     float32 // ±3.4e38 (~6-7 decimal digits precision)
     float64 // ±1.8e308 (~15-17 decimal digits precision)
     ```
   - Complex Numbers:
     ```go
     complex64  // Complex numbers with float32 parts
     complex128 // Complex numbers with float64 parts
     ```

3. **UserInput** - Working with user input:
   ```go
   // Basic input
   var name string
   fmt.Scan(&name)

   // Formatted input
   var age int
   fmt.Scanf("%d", &age)

   // String input with spaces
   reader := bufio.NewReader(os.Stdin)
   text, _ := reader.ReadString('\n')
   ```

4. **Conversions** - Type conversions:
   ```go
   // Basic type conversion
   i := 42
   f := float64(i)
   
   // String conversions
   str := strconv.Itoa(i)      // int to string
   num, err := strconv.Atoi(str) // string to int
   ```

5. **MyMaths** - Mathematical operations
6. **MyTimes** - Time and date handling:
   ```go
   now := time.Now()
   formatted := now.Format("2006-01-02 15:04:05")
   ```

### Intermediate Concepts
7. **MyPointers** - Understanding pointers:
   ```go
   var p *int
   i := 42
   p = &i    // p points to i
   *p = 21   // i is now 21
   ```

8. **Array** - Working with arrays:
   ```go
   var numbers [5]int
   names := [3]string{"John", "Jane", "Joe"}
   matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}
   ```

9. **Slices** - Dynamic arrays:
   ```go
   slice := make([]int, 5, 10)  // length 5, capacity 10
   slice = append(slice, 1, 2, 3)
   ```

10. **Maps** - Key-value pairs:
    ```go
    m := make(map[string]int)
    m["key"] = 42
    value, exists := m["key"]
    delete(m, "key")
    ```

11. **Structs** - Custom data types:
    ```go
    type Person struct {
        Name    string
        Age     int
        Address Address
    }
    ```

### Control Flow
12. **Conditions** - Control flow:
    ```go
    if x > 0 {
        // positive
    } else if x < 0 {
        // negative
    } else {
        // zero
    }
    ```

13. **SwitchCase** - Switch statements:
    ```go
    switch value := expr.(type) {
    case int:
        fmt.Println("Integer")
    case string:
        fmt.Println("String")
    default:
        fmt.Println("Unknown type")
    }
    ```

14. **Loops** - Iteration:
    ```go
    // Basic for loop
    for i := 0; i < 10; i++ {}
    
    // Range loop
    for index, value := range slice {}
    
    // Infinite loop
    for {}
    ```

### Functions and Methods
15. **Functions** - Various function types:
    ```go
    // Basic function
    func add(x, y int) int {
        return x + y
    }

    // Multiple returns
    func divide(x, y float64) (float64, error) {
        if y == 0 {
            return 0, errors.New("division by zero")
        }
        return x / y, nil
    }
    ```

16. **Methods** - Object-oriented programming:
    ```go
    type Rectangle struct {
        width, height float64
    }

    func (r Rectangle) Area() float64 {
        return r.width * r.height
    }
    ```

17. **Defer** - Deferred execution:
    ```go
    func example() {
        defer fmt.Println("cleanup")
        // main logic
    }
    ```

### File Operations and Web
18. **HandlingFiles** - File operations:
    ```go
    // Reading file
    data, err := ioutil.ReadFile("file.txt")
    
    // Writing file
    err = ioutil.WriteFile("file.txt", data, 0644)
    ```

19. **WebReqHandel** - Web programming:
    ```go
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    http.ListenAndServe(":8080", nil)
    ```

## Additional Topics

### Error Handling
```go
// Custom error type
type CustomError struct {
    Code    int
    Message string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// Error handling pattern
if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}
```

### Concurrency
```go
// Goroutine
go func() {
    // concurrent code
}()

// Channel operations
ch := make(chan int, 100)
ch <- 42      // send
value := <-ch // receive
close(ch)     // close
```

## Prerequisites
- Basic programming knowledge
- Go installed on your system
- Text editor or IDE (VSCode recommended)

## Getting Started
1. Clone this repository
2. Navigate to each folder sequentially
3. Read the accompanying documentation
4. Complete the exercises in each section

## Additional Resources
- [Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Go Playground](https://play.golang.org/)

## Contributing
Feel free to contribute to this course by:
- Reporting issues
- Suggesting improvements
- Adding more examples
- Fixing errors

## License
This course is available under the MIT License.