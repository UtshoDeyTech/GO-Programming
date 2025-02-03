// 1. Basic Functions and Advanced Concepts
package main

import "fmt"

func main() {
	// Basic function calls
	fmt.Println("\n1. Basic Functions:")
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()
	result := adder(3, 5)
	fmt.Println("Result is", result)
	proResult := proAdder(2, 3, 4, 5, 6)
	fmt.Println("Pro Result is", proResult)

	// 2. Multiple Return Values
	fmt.Println("\n2. Multiple Return Values:")
	sum, diff, prod, div := calculator(10, 5)
	fmt.Printf("Calculator results: sum=%d, diff=%d, prod=%d, div=%.2f\n",
		sum, diff, prod, div)

	// 3. Named Return Values
	fmt.Println("\n3. Named Return Values:")
	area, perim := rectangle(5.0, 3.0)
	fmt.Printf("Rectangle: area=%.2f, perimeter=%.2f\n", area, perim)

	// 4. Higher-Order Function
	fmt.Println("\n4. Higher-Order Function:")
	result = operate(4, 5, multiply)
	fmt.Printf("Higher-order function result: %d\n", result)

	// 5. Closure
	fmt.Println("\n5. Closure Example:")
	increment := counter()
	fmt.Println("First call:", increment())
	fmt.Println("Second call:", increment())

	// 6. Method
	fmt.Println("\n6. Method Example:")
	c := Circle{radius: 5}
	fmt.Printf("Circle area: %.2f\n", c.area())

	// 7. Error Handling
	fmt.Println("\n7. Error Handling:")
	if result, err := divide(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Division result: %.2f\n", result)
	}

	// Try division by zero
	if _, err := divide(10, 0); err != nil {
		fmt.Println("Error case:", err)
	}

	// 8. Variadic with Different Types
	fmt.Println("\n8. Variadic Function with Different Types:")
	printValues("hello", 42, true, 3.14)
}

// Basic greeting function
func greeter() {
	fmt.Println("Hello, I am utsho dey")
}

// Another greeting function
func greeterTwo() {
	fmt.Println("Hello again, I am utsho dey")
}

// Simple addition function
func adder(val1 int, val2 int) int {
	return val1 + val2
}

// Variadic function for adding multiple numbers
func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

// 2. Function with multiple return values
func calculator(a, b int) (int, int, int, float64) {
	sum := a + b
	diff := a - b
	product := a * b
	division := float64(a) / float64(b)
	return sum, diff, product, division
}

// 3. Function with named return values
func rectangle(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return
}

// 4. Higher-order function
func operate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func multiply(x, y int) int {
	return x * y
}

// 5. Closure
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 6. Method definition
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// 7. Function with error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// 8. Variadic function with different types
func printValues(values ...interface{}) {
	for _, value := range values {
		fmt.Printf("%v ", value)
	}
	fmt.Println()
}
