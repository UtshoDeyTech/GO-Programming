package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// Original crypto random number
	fmt.Println("\n=== Crypto Random Numbers ===")
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("Crypto random number:", myRandomNumber)

	// Generate multiple crypto random numbers
	fmt.Println("\n=== Multiple Crypto Random Numbers ===")
	for i := 0; i < 5; i++ {
		randNum, _ := rand.Int(rand.Reader, big.NewInt(100))
		fmt.Printf("Random number %d: %v\n", i+1, randNum)
	}

	// Basic Math Operations
	fmt.Println("\n=== Basic Math Operations ===")
	x, y := 15, 4
	fmt.Printf("x = %d, y = %d\n", x, y)
	fmt.Printf("Addition: %d\n", x+y)
	fmt.Printf("Subtraction: %d\n", x-y)
	fmt.Printf("Multiplication: %d\n", x*y)
	fmt.Printf("Division: %.2f\n", float64(x)/float64(y))
	fmt.Printf("Modulus: %d\n", x%y)

	// Math Package Functions
	fmt.Println("\n=== Math Package Functions ===")
	num := 16.0
	fmt.Printf("Square root of %.0f: %.2f\n", num, math.Sqrt(num))
	fmt.Printf("Power of %.0f^2: %.0f\n", num, math.Pow(num, 2))
	fmt.Printf("Ceil of 3.14: %.0f\n", math.Ceil(3.14))
	fmt.Printf("Floor of 3.14: %.0f\n", math.Floor(3.14))
	fmt.Printf("Round of 3.51: %.0f\n", math.Round(3.51))

	// Trigonometric Functions
	fmt.Println("\n=== Trigonometric Functions ===")
	angle := math.Pi / 4 // 45 degrees
	fmt.Printf("Sin(π/4): %.2f\n", math.Sin(angle))
	fmt.Printf("Cos(π/4): %.2f\n", math.Cos(angle))
	fmt.Printf("Tan(π/4): %.2f\n", math.Tan(angle))

	// Constants from math package
	fmt.Println("\n=== Mathematical Constants ===")
	fmt.Printf("Pi: %.5f\n", math.Pi)
	fmt.Printf("E: %.5f\n", math.E)

	// Working with Big Numbers
	fmt.Println("\n=== Big Numbers ===")
	// Creating big integers
	bigNum1 := new(big.Int)
	bigNum1.SetString("12345678901234567890", 10)
	bigNum2 := new(big.Int)
	bigNum2.SetString("98765432109876543210", 10)

	// Big number operations
	result := new(big.Int)
	// Addition
	result.Add(bigNum1, bigNum2)
	fmt.Println("Big number addition:", result)
	// Multiplication
	result.Mul(bigNum1, bigNum2)
	fmt.Println("Big number multiplication:", result)

	// Working with Big Floats
	fmt.Println("\n=== Big Floats ===")
	pi := new(big.Float)
	pi.SetPrec(100)
	pi.SetString("3.14159265358979323846264338327950288419716939937510582097494459")

	radius := new(big.Float).SetFloat64(2.0)
	area := new(big.Float).Mul(pi, new(big.Float).Mul(radius, radius))
	fmt.Printf("Area of circle with radius 2: %.10f\n", area)

	// Absolute Values
	fmt.Println("\n=== Absolute Values ===")
	fmt.Printf("Absolute of -15.5: %.2f\n", math.Abs(-15.5))
	fmt.Printf("Absolute of -42: %d\n", abs(-42))

	// Min and Max
	fmt.Println("\n=== Min and Max ===")
	numbers := []float64{23.5, 11.2, 42.8, 19.9}
	fmt.Printf("Min value: %.2f\n", findMin(numbers))
	fmt.Printf("Max value: %.2f\n", findMax(numbers))

	// Calculate factorial using big.Int
	fmt.Println("\n=== Factorial Calculation ===")
	n := 20
	fmt.Printf("Factorial of %d: %v\n", n, factorial(n))
}

// Helper function for integer absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Helper function to find minimum value in slice
func findMin(numbers []float64) float64 {
	min := numbers[0]
	for _, num := range numbers {
		min = math.Min(min, num)
	}
	return min
}

// Helper function to find maximum value in slice
func findMax(numbers []float64) float64 {
	max := numbers[0]
	for _, num := range numbers {
		max = math.Max(max, num)
	}
	return max
}

// Calculate factorial using big.Int
func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
