package main

import (
	"fmt"
)

// Function that takes a pointer parameter
func updateValue(ptr *int) {
	*ptr = *ptr * 2
}

// Function that returns a pointer
func createPointer() *int {
	value := 42
	return &value
}

// Struct with pointers
type Person struct {
	name  *string
	age   *int
	score *float64
}

func main() {
	fmt.Println("Welcome to pointers")

	// Your original code
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("\n=== Original Code ===")
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value stored at pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is", myNumber)

	// Basic Pointer Operations
	fmt.Println("\n=== Basic Pointer Operations ===")
	x := 100
	ptr1 := &x
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", ptr1)
	fmt.Printf("Value at pointer: %d\n", *ptr1)

	// Changing value through pointer
	*ptr1 = 200
	fmt.Printf("New value of x: %d\n", x)

	// Pointer to Pointer
	fmt.Println("\n=== Pointer to Pointer ===")
	var ptr2 **int = &ptr1
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Value using ptr1: %d\n", *ptr1)
	fmt.Printf("Value using ptr2: %d\n", **ptr2)

	// Modify value using double pointer
	**ptr2 = 300
	fmt.Printf("New value of x: %d\n", x)

	// Array of Pointers
	fmt.Println("\n=== Array of Pointers ===")
	a, b, c := 1, 2, 3
	pointerArray := []*int{&a, &b, &c}
	for i, ptr := range pointerArray {
		fmt.Printf("Pointer %d points to value: %d\n", i, *ptr)
	}

	// Using pointer function
	fmt.Println("\n=== Pointer Function ===")
	value := 5
	fmt.Printf("Before update: %d\n", value)
	updateValue(&value)
	fmt.Printf("After update: %d\n", value)

	// Getting pointer from function
	fmt.Println("\n=== Function Returning Pointer ===")
	newPtr := createPointer()
	fmt.Printf("Value from pointer function: %d\n", *newPtr)

	// Struct with Pointers
	fmt.Println("\n=== Struct with Pointers ===")
	name := "John"
	age := 30
	score := 85.5

	person := Person{
		name:  &name,
		age:   &age,
		score: &score,
	}

	fmt.Printf("Name: %s\n", *person.name)
	fmt.Printf("Age: %d\n", *person.age)
	fmt.Printf("Score: %.1f\n", *person.score)

	// Modifying struct values through pointers
	*person.age = 31
	fmt.Printf("Updated age: %d\n", *person.age)

	// Nil Pointer Example
	fmt.Println("\n=== Nil Pointer Example ===")
	var nilPtr *int
	maybePtr := createMaybeNilPointer() // Function that might return nil

	fmt.Println("Working with potentially nil pointers:")
	showPointerValue(nilPtr)    // Will show "Pointer is nil"
	showPointerValue(maybePtr)  // Will show value if not nil
	showPointerValue(&myNumber) // Will show actual value

	// Pointer Arithmetic (not directly supported in Go)
	fmt.Println("\n=== Slice Instead of Pointer Arithmetic ===")
	numbers := []int{1, 2, 3, 4, 5}
	// Instead of pointer arithmetic, use slice operations
	fmt.Println("First element:", numbers[0])
	fmt.Println("Rest of slice:", numbers[1:])

	// Memory Management Example
	fmt.Println("\n=== Memory Management ===")
	someFunc()
}

func someFunc() {
	localVar := 42
	ptr := &localVar
	fmt.Printf("Local variable value: %d\n", *ptr)
	fmt.Printf("Local variable address: %p\n", ptr)
}

// Helper function to demonstrate nil pointer handling
func createMaybeNilPointer() *int {
	// Simulating a function that might return nil
	return nil
}

// Helper function to safely show pointer value
func showPointerValue(ptr *int) {
	if ptr == nil {
		fmt.Println("Pointer is nil")
		return
	}
	fmt.Printf("Pointer value: %d\n", *ptr)
}
