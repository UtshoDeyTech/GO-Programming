package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("If else in golang")

	// Original login count example
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "It's exactly 10"
	}
	fmt.Println("Login result:", result)

	// Even/Odd check
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// Initialization with if statement
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Number is more than 10")
	}

	// Multiple conditions using &&
	age := 18
	hasID := true
	if age >= 18 && hasID {
		fmt.Println("Can enter the venue")
	} else {
		fmt.Println("Cannot enter the venue")
	}

	// Random number with multiple conditions
	rand.Seed(time.Now().UnixNano())
	score := rand.Intn(100) + 1
	fmt.Printf("Score is: %d\n", score)

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Nested if statements
	temp := 25
	time := 14 // 24-hour format

	if temp > 20 {
		if time >= 12 && time <= 18 {
			fmt.Println("Perfect time for swimming")
		} else {
			fmt.Println("Good temperature but not the right time")
		}
	} else {
		fmt.Println("Too cold for swimming")
	}

	// Using OR operator ||
	isWeekend := true
	isHoliday := false
	if isWeekend || isHoliday {
		fmt.Println("You can sleep late")
	} else {
		fmt.Println("Wake up early!")
	}

	// Error checking pattern (common in Go)
	if value, err := someFunction(); err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Value is:", value)
	}
}

// Helper function for error checking example
func someFunction() (int, error) {
	// Simulating a function that might return an error
	if rand.Float64() < 0.5 {
		return 0, fmt.Errorf("something went wrong")
	}
	return 42, nil
}
