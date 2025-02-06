package main

import (
	"fmt"
	"time"
)

type Person struct {
	name  *string
	age   *int
	score *float64
}

type Node struct {
	data int
	next *Node
	prev *Node
}

// Helper function for pointer checking
func printPointerValue(ptr *int) {
	if ptr == nil {
		fmt.Println("Pointer is nil - cannot dereference")
		return
	}
	fmt.Printf("Pointer value: %d\n", *ptr)
}

func main() {
	// Initialize and clear screen
	fmt.Print("\033[H\033[2J")
	fmt.Println("\033[1;36m=== Pointer Operations in Go ===\033[0m")
	fmt.Println("\033[1;33mLearning about memory addresses and references\033[0m")
	time.Sleep(1 * time.Second)

	// 1. Basic Pointer Concepts
	fmt.Println("\n\033[1;32m1. BASIC POINTER CONCEPTS\033[0m")
	fmt.Println("------------------------")

	// Initialize a variable
	myNumber := 42

	// Create a pointer to the variable
	myPointer := &myNumber

	fmt.Printf("Variable Value:    %d\n", myNumber)
	fmt.Printf("Variable Address:  %p\n", myPointer)
	fmt.Printf("Pointer Value:     %p\n", myPointer)
	fmt.Printf("Dereferenced Value: %d\n", *myPointer)

	// Modify value through pointer
	*myPointer = 100
	fmt.Printf("Modified Value:    %d\n", myNumber)

	// 2. Multiple Pointers
	fmt.Println("\n\033[1;32m2. MULTIPLE POINTERS\033[0m")
	fmt.Println("-------------------")

	x := 25
	pointer1 := &x
	pointer2 := &x

	fmt.Printf("Original Value: %d\n", x)
	*pointer1 = 50
	fmt.Printf("After pointer1 modification: %d\n", x)
	*pointer2 = 75
	fmt.Printf("After pointer2 modification: %d\n", x)

	// 3. Pointer to Pointer
	fmt.Println("\n\033[1;32m3. POINTER TO POINTER\033[0m")
	fmt.Println("-------------------")

	value := 300
	ptr1 := &value
	ptr2 := &ptr1

	fmt.Printf("Value: %d\n", value)
	fmt.Printf("Pointer1 -> Value: %d\n", *ptr1)
	fmt.Printf("Pointer2 -> Pointer1 -> Value: %d\n", **ptr2)

	// Modify through double pointer
	**ptr2 = 400
	fmt.Printf("Modified Value: %d\n", value)

	// 4. Pointers with Arrays and Slices
	fmt.Println("\n\033[1;32m4. POINTERS WITH ARRAYS AND SLICES\033[0m")
	fmt.Println("--------------------------------")

	// Array of integers
	numbers := [3]int{10, 20, 30}

	// Array of pointers
	var pointerArray [3]*int
	for i := range numbers {
		pointerArray[i] = &numbers[i]
	}

	// Print and modify through pointers
	for i, ptr := range pointerArray {
		fmt.Printf("Original value at index %d: %d\n", i, *ptr)
		*ptr *= 2
		fmt.Printf("Modified value at index %d: %d\n", i, numbers[i])
	}

	// 5. Struct with Pointers
	fmt.Println("\n\033[1;32m5. STRUCT WITH POINTERS\033[0m")
	fmt.Println("----------------------")

	name := "Alice"
	age := 25
	score := 95.5

	// Create struct with pointers
	person := Person{
		name:  &name,
		age:   &age,
		score: &score,
	}

	// Access and modify through pointers
	fmt.Printf("Name: %s\n", *person.name)
	fmt.Printf("Age: %d\n", *person.age)
	fmt.Printf("Score: %.1f\n", *person.score)

	*person.age = 26
	*person.score = 97.5
	fmt.Printf("\nUpdated Age: %d\n", *person.age)
	fmt.Printf("Updated Score: %.1f\n", *person.score)

	// 6. Nil Pointer Handling
	fmt.Println("\n\033[1;32m6. NIL POINTER HANDLING\033[0m")
	fmt.Println("----------------------")

	// Example 1: Working with nil pointers
	var nilPtr *int
	fmt.Printf("Nil pointer value: %v\n", nilPtr)

	// Safe pointer handling
	safeValue := 42
	safePtr := &safeValue

	// Example of checking both nil and non-nil pointers
	printPointerValue(nilPtr)  // Will show "Pointer is nil"
	printPointerValue(safePtr) // Will show the value

	// Initialize previously nil pointer
	value = 500
	nilPtr = &value
	fmt.Printf("No longer nil, value is: %d\n", *nilPtr)

	// 7. Linked List with Pointers
	fmt.Println("\n\033[1;32m7. LINKED LIST WITH POINTERS\033[0m")
	fmt.Println("---------------------------")

	// Create nodes
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}

	// Link nodes
	node1.next = node2
	node2.prev = node1
	node2.next = node3
	node3.prev = node2

	// Traverse forward
	fmt.Println("Forward traversal:")
	current := node1
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")

	// Traverse backward
	fmt.Println("Backward traversal:")
	current = node3
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.prev
	}
	fmt.Println("nil")

	// 8. Memory Address Exploration
	fmt.Println("\n\033[1;32m8. MEMORY ADDRESS EXPLORATION\033[0m")
	fmt.Println("-----------------------------")

	var (
		intVar    int     = 42
		floatVar  float64 = 3.14
		stringVar string  = "Hello"
		boolVar   bool    = true
	)

	fmt.Printf("Integer Address:  %p\n", &intVar)
	fmt.Printf("Float Address:    %p\n", &floatVar)
	fmt.Printf("String Address:   %p\n", &stringVar)
	fmt.Printf("Boolean Address:  %p\n", &boolVar)

	// 9. Pointer Value Operations
	fmt.Println("\n\033[1;32m9. POINTER VALUE OPERATIONS\033[0m")
	fmt.Println("--------------------------")

	integer := 42
	intPtr := &integer

	// Create new value from pointer
	newInt := *intPtr
	fmt.Printf("Original Value:     %d\n", integer)
	fmt.Printf("Pointer Address:    %p\n", intPtr)
	fmt.Printf("Dereferenced Value: %d\n", newInt)

	// Modify through pointer
	*intPtr += 10
	fmt.Printf("Modified Value:     %d\n", integer)

	// Program completion
	fmt.Println("\n\033[1;36m=== Program Complete ===\033[0m")
	fmt.Println("\033[1;33mPress Enter to exit...\033[0m")
	fmt.Scanln()
}
