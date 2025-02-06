package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
	"unsafe"
)

// Public constants and variables
const LoginToken string = "sbfjkskds"

var jwtToken int = 30000

// Additional constants
const (
	Pi             = 3.14159
	MaxConnections = 100
	DatabaseURL    = "localhost:5432"
)

// Days of week using iota
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Size constants
const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
)

func main() {
	fmt.Println("=== Basic Variable Declarations ===")

	// Using public variable
	fmt.Printf("JWT Token: %v\n", jwtToken)

	// Basic string variable
	var username string = "utsho"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	// Boolean variable
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	// Integer variable
	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	// Float variables
	var smallFloat float32 = 255.387478242492378
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	var longFloat float64 = 255.387478242492378
	fmt.Println(longFloat)
	fmt.Printf("Variable is of type: %T\n", longFloat)

	// Default value
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	// Implicit type
	var website = "utshodey.tech"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T\n", website)

	// Short declaration
	numberofUser := 300000.0
	fmt.Println(numberofUser)
	fmt.Printf("Variable is of type: %T\n", numberofUser)

	fmt.Println("\n=== Numeric Types ===")

	// Different integer types
	var (
		age         uint8  = 25
		temperature int8   = -10
		population  uint16 = 65000
		distance    int16  = -30000
		salary      uint32 = 50000
		worldPop    int64  = 7800000000
	)

	fmt.Printf("age (uint8): %v\n", age)
	fmt.Printf("temperature (int8): %v\n", temperature)
	fmt.Printf("population (uint16): %v\n", population)
	fmt.Printf("distance (int16): %v\n", distance)
	fmt.Printf("salary (uint32): %v\n", salary)
	fmt.Printf("world population (int64): %v\n", worldPop)

	// Complex numbers
	var (
		c64  complex64  = 5 + 10i
		c128 complex128 = 10 + 15i
	)
	fmt.Printf("complex64: %v\n", c64)
	fmt.Printf("complex128: %v\n", c128)

	fmt.Println("\n=== String Operations ===")

	var nameOne string = "Utsho"
	nameTwo := "Dey"
	fullName := nameOne + " " + nameTwo

	fmt.Printf("Full name: %s\n", fullName)
	fmt.Printf("Length of full name: %d\n", len(fullName))
	fmt.Printf("First character of name: %c\n", fullName[0])

	multiLineStr := `This is a
    multiline string
    that preserves formatting
    and line breaks`
	fmt.Println("Multiline string:")
	fmt.Println(multiLineStr)

	fmt.Println("\n=== Constants ===")

	fmt.Printf("Pi: %v\n", Pi)
	fmt.Printf("Max Connections: %v\n", MaxConnections)
	fmt.Printf("Database URL: %v\n", DatabaseURL)
	fmt.Printf("Days: Sunday=%v, Monday=%v, Saturday=%v\n", Sunday, Monday, Saturday)
	fmt.Printf("Size constants: KB=%d, MB=%d, GB=%d, TB=%d\n", KB, MB, GB, TB)

	fmt.Println("\n=== Type Conversion ===")

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Printf("int: %v, float64: %v, uint: %v\n", i, f, u)

	// String conversion
	str := "123"
	if num, err := strconv.Atoi(str); err == nil {
		fmt.Printf("String to int conversion: %d\n", num)
	}

	fmt.Println("\n=== Memory Sizes ===")

	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(int(0)))
	fmt.Printf("Size of float64: %d bytes\n", unsafe.Sizeof(float64(0)))
	fmt.Printf("Size of string: %d bytes\n", unsafe.Sizeof(string("")))
	fmt.Printf("Size of bool: %d bytes\n", unsafe.Sizeof(bool(false)))

	fmt.Println("\n=== Math Operations ===")

	var (
		intNum     int     = 42
		floatNum   float64 = 42.5
		complexNum         = 3 + 4i
	)

	fmt.Printf("Square root of float: %f\n", math.Sqrt(floatNum))
	fmt.Printf("Complex number magnitude: %f\n", cmplx.Abs(complexNum))
	fmt.Printf("Integer power: %d\n", intNum*intNum)

	// Pointer example
	x := 42
	ptr := &x
	fmt.Printf("Value: %d, Pointer: %p, Dereferenced value: %d\n", x, ptr, *ptr)
}
