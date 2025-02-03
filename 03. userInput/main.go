package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	// Basic string input (your original code)
	fmt.Println("\n=== Basic String Input ===")
	fmt.Println("Enter the rating for our Pizza:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank for rating:", input)
	fmt.Printf("Type of the rating is %T\n", input)

	// Proper error handling
	fmt.Println("\n=== Error Handling ===")
	fmt.Println("Enter another rating:")
	rating, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred:", err)
	} else {
		fmt.Println("Thank you for rating:", rating)
	}

	// Converting string to number
	fmt.Println("\n=== Number Conversion ===")
	fmt.Println("Enter a number between 1-5:")
	numInput, _ := reader.ReadString('\n')
	// Trim the newline and any spaces
	numInput = strings.TrimSpace(numInput)

	num, err := strconv.ParseFloat(numInput, 64)
	if err != nil {
		fmt.Println("Error converting to number:", err)
	} else {
		fmt.Printf("Number entered: %v, Type: %T\n", num, num)
		// Basic input validation
		if num >= 1 && num <= 5 {
			fmt.Println("Valid rating!")
		} else {
			fmt.Println("Please enter a rating between 1 and 5")
		}
	}

	// Multiple inputs on one line
	fmt.Println("\n=== Multiple Inputs ===")
	fmt.Println("Enter your first and last name (separated by space):")
	nameInput, _ := reader.ReadString('\n')
	nameInput = strings.TrimSpace(nameInput)
	names := strings.Split(nameInput, " ")
	if len(names) >= 2 {
		fmt.Printf("First Name: %s\nLast Name: %s\n", names[0], names[1])
	}

	// Password input (simple example)
	fmt.Println("\n=== Password Input ===")
	fmt.Println("Enter your password (it will be visible):")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)
	if len(password) < 8 {
		fmt.Println("Password should be at least 8 characters long")
	} else {
		fmt.Println("Password accepted!")
	}

	// Menu-driven input
	fmt.Println("\n=== Menu Selection ===")
	fmt.Println("Select an option:")
	fmt.Println("1. View profile")
	fmt.Println("2. Edit profile")
	fmt.Println("3. Exit")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("Viewing profile...")
	case "2":
		fmt.Println("Editing profile...")
	case "3":
		fmt.Println("Exiting...")
	default:
		fmt.Println("Invalid option selected")
	}

	// Continuous input until a condition is met
	fmt.Println("\n=== Continuous Input ===")
	fmt.Println("Keep entering numbers to add (enter 'sum' to calculate total):")

	var numbers []float64
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "sum" {
			// Calculate sum
			var sum float64
			for _, num := range numbers {
				sum += num
			}
			fmt.Printf("Total sum: %.2f\n", sum)
			break
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Please enter a valid number or 'sum'")
			continue
		}
		numbers = append(numbers, num)
	}

	// File input example
	fmt.Println("\n=== File Input ===")
	fmt.Println("Enter a filename to create:")
	filename, _ := reader.ReadString('\n')
	filename = strings.TrimSpace(filename)

	fmt.Println("Enter content for the file:")
	content, _ := reader.ReadString('\n')

	// Writing to file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
	} else {
		defer file.Close()
		_, err := file.WriteString(content)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		} else {
			fmt.Println("Successfully wrote to file!")
		}
	}
}
