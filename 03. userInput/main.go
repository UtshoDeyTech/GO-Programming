package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Initialize reader for user input
	reader := bufio.NewReader(os.Stdin)

	// Clear screen and show welcome message
	fmt.Print("\033[H\033[2J") // ANSI escape code to clear screen
	fmt.Println("=== Welcome to the Interactive Go Program ===")
	time.Sleep(1 * time.Second) // Add a small delay for better UX

	// 1. Basic String Input with Validation
	fmt.Println("\n1. USER INFORMATION")
	fmt.Println("-----------------")

	// Name input with validation
	var firstName, lastName string
	for {
		fmt.Print("Enter your first and last name (e.g., John Doe): ")
		nameInput, _ := reader.ReadString('\n')
		nameInput = strings.TrimSpace(nameInput)
		names := strings.Split(nameInput, " ")

		if len(names) < 2 {
			fmt.Println("\033[31mPlease enter both first and last name!\033[0m") // Red color
			continue
		}
		firstName, lastName = names[0], names[1]
		break
	}

	// 2. Age Input with Validation
	var age int
	for {
		fmt.Print("Enter your age (18-100): ")
		ageInput, _ := reader.ReadString('\n')
		ageInput = strings.TrimSpace(ageInput)

		ageNum, err := strconv.Atoi(ageInput)
		if err != nil || ageNum < 18 || ageNum > 100 {
			fmt.Println("\033[31mPlease enter a valid age between 18 and 100!\033[0m")
			continue
		}
		age = ageNum
		break
	}

	// 3. Password Creation with Strength Check
	var password string
	for {
		fmt.Print("Create a password (min 8 chars, must include numbers): ")
		passInput, _ := reader.ReadString('\n')
		password = strings.TrimSpace(passInput)

		hasNumber := false
		for _, char := range password {
			if char >= '0' && char <= '9' {
				hasNumber = true
				break
			}
		}

		if len(password) < 8 || !hasNumber {
			fmt.Println("\033[31mPassword must be at least 8 characters and contain numbers!\033[0m")
			continue
		}
		break
	}

	// 4. Rating System with Float Values
	var rating float64
	for {
		fmt.Print("Rate your experience (1.0-5.0): ")
		ratingInput, _ := reader.ReadString('\n')
		ratingInput = strings.TrimSpace(ratingInput)

		ratingNum, err := strconv.ParseFloat(ratingInput, 64)
		if err != nil || ratingNum < 1.0 || ratingNum > 5.0 {
			fmt.Println("\033[31mPlease enter a valid rating between 1.0 and 5.0!\033[0m")
			continue
		}
		rating = ratingNum
		break
	}

	// 5. Simple File Operations
	fmt.Println("\n2. FILE OPERATIONS")
	fmt.Println("----------------")

	// Create a user profile file
	filename := fmt.Sprintf("%s_%s_profile.txt", firstName, lastName)
	content := fmt.Sprintf("User Profile:\nName: %s %s\nAge: %d\nRating: %.1f\nCreated: %s",
		firstName, lastName, age, rating, time.Now().Format("2006-01-02 15:04:05"))

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("\033[31mError creating file: %v\033[0m\n", err)
	} else {
		defer file.Close()
		_, err := file.WriteString(content)
		if err != nil {
			fmt.Printf("\033[31mError writing to file: %v\033[0m\n", err)
		} else {
			fmt.Printf("\033[32mProfile successfully saved to %s!\033[0m\n", filename)
		}
	}

	// 6. Calculator Feature
	fmt.Println("\n3. SIMPLE CALCULATOR")
	fmt.Println("-----------------")
	fmt.Println("Enter numbers to add (type 'done' to finish):")

	var numbers []float64
	var sum float64

	for {
		fmt.Print("Enter number (or 'done'): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "done" {
			break
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("\033[31mPlease enter a valid number!\033[0m")
			continue
		}

		numbers = append(numbers, num)
		sum += num
		fmt.Printf("\033[34mRunning total: %.2f\033[0m\n", sum)
	}

	fmt.Printf("\n\033[32mFinal sum: %.2f\033[0m\n", sum)
	if len(numbers) > 0 {
		fmt.Printf("\033[32mAverage: %.2f\033[0m\n", sum/float64(len(numbers)))
	}

	// 7. Program Exit
	fmt.Println("\n=== Program Complete ===")
	fmt.Print("Press Enter to exit...")
	reader.ReadString('\n')
}
