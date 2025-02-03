package main

import "fmt"

// Public constants and variables (package level)
const LoginToken string = "sbfjkskds" // public
var jwtToken int = 30000              // public

// Additional constants
const (
	Pi             = 3.14159
	MaxConnections = 100
	DatabaseURL    = "localhost:5432"
)

// Using iota for enumerated constants
const (
	Sunday = iota // 0
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Type aliases
type UserID uint64
type Temperature float64
type Status int

func main() {

	// Using jwtToken to avoid unused variable warning
	fmt.Printf("JWT Token: %v\n", jwtToken)

	// Your original code
	var username string = "utsho"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float32 = 255.387478242492378
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	var longFloat float64 = 255.387478242492378
	fmt.Println(longFloat)
	fmt.Printf("Variable is of type : %T \n", longFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// implicit type
	var website = "utshodey.tech"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)

	// no var style
	numberofUser := 300000.0
	fmt.Println(numberofUser)
	fmt.Printf("Variable is of type : %T \n", website)

	// Numeric Types Examples
	fmt.Println("\n=== Numeric Types ===")
	var (
		age         uint8  = 25         // 0 to 255
		temperature int8   = -10        // -128 to 127
		population  uint16 = 65000      // 0 to 65535
		distance    int16  = -30000     // -32768 to 32767
		salary      uint32 = 50000      // 0 to 4294967295
		worldPop    int64  = 7800000000 // -9223372036854775808 to 9223372036854775807
	)

	fmt.Printf("age (uint8): %v\n", age)
	fmt.Printf("temperature (int8): %v\n", temperature)
	fmt.Printf("population (uint16): %v\n", population)
	fmt.Printf("distance (int16): %v\n", distance)
	fmt.Printf("salary (uint32): %v\n", salary)
	fmt.Printf("world population (int64): %v\n", worldPop)

	// Complex numbers
	var (
		c64  complex64  = 5 + 10i  // 32-bit real and imaginary parts
		c128 complex128 = 10 + 15i // 64-bit real and imaginary parts
	)
	fmt.Printf("complex64: %v\n", c64)
	fmt.Printf("complex128: %v\n", c128)

	// String Operations Examples
	fmt.Println("\n=== String Operations ===")
	var nameOne string = "Utsho"
	nameTwo := "Dey"

	// String concatenation
	fullName := nameOne + " " + nameTwo
	fmt.Printf("Full name: %s\n", fullName)

	// String length
	fmt.Printf("Length of full name: %d\n", len(fullName))

	// Raw string literal (preserves formatting)
	multiLineStr := `This is a
    multiline string
    that preserves formatting
    and line breaks`
	fmt.Println("Multiline string:")
	fmt.Println(multiLineStr)

	// String indexing
	fmt.Printf("First character of name: %c\n", fullName[0])

	// Constants Examples
	fmt.Println("\n=== Constants ===")
	fmt.Printf("Pi: %v\n", Pi)
	fmt.Printf("Max Connections: %v\n", MaxConnections)
	fmt.Printf("Database URL: %v\n", DatabaseURL)

	fmt.Println("\nDays of week:")
	fmt.Printf("Sunday: %v\n", Sunday)
	fmt.Printf("Monday: %v\n", Monday)
	fmt.Printf("Saturday: %v\n", Saturday)

	// Type Conversion Examples
	fmt.Println("\n=== Type Conversion ===")

	// Basic type conversion
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("int: %v, float64: %v, uint: %v\n", i, f, u)

	// Using type aliases
	var userId UserID = 12345
	var temp Temperature = 36.6
	var status Status = 1

	fmt.Printf("UserID: %v (%T)\n", userId, userId)
	fmt.Printf("Temperature: %v (%T)\n", temp, temp)
	fmt.Printf("Status: %v (%T)\n", status, status)

	// Multiple variable declarations
	fmt.Println("\n=== Multiple Variable Declarations ===")
	var (
		name    string  = "John"
		age2    int     = 25
		height  float64 = 1.75
		isAdmin bool    = false
	)
	fmt.Printf("Name: %v, Age: %v, Height: %v, Admin: %v\n", name, age2, height, isAdmin)

	// Zero values
	fmt.Println("\n=== Zero Values ===")
	var (
		zeroInt    int
		zeroFloat  float64
		zeroString string
		zeroBool   bool
		zeroPtr    *int
	)
	fmt.Printf("Zero values - Int: %v, Float: %v, String: %q, Bool: %v, Pointer: %v\n",
		zeroInt, zeroFloat, zeroString, zeroBool, zeroPtr)
}
