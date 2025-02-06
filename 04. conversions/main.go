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
	// Initialize reader and clear screen
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[H\033[2J") // Clear screen

	// Welcome message with colored output
	fmt.Println("\033[1;36m=== Go Type Conversion Practice ===\033[0m")
	fmt.Println("\033[1;33mLet's learn about different type conversions in Go!\033[0m")
	time.Sleep(1 * time.Second)

	// 1. Basic String to Number Conversions
	fmt.Println("\n\033[1;32m1. BASIC NUMBER CONVERSIONS\033[0m")
	fmt.Println("---------------------------")

	// Float conversion with validation
	fmt.Print("Enter a decimal number (e.g., 3.14): ")
	input, _ := reader.ReadString('\n')
	numFloat, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Printf("\033[31mError: Invalid decimal number (%v)\033[0m\n", err)
	} else {
		fmt.Printf("As float64: %f\nRounded: %.2f\nScientific: %e\n",
			numFloat, numFloat, numFloat)
	}

	// Integer conversion with validation
	fmt.Print("\nEnter a whole number: ")
	input, _ = reader.ReadString('\n')
	numInt, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Printf("\033[31mError: Invalid integer (%v)\033[0m\n", err)
	} else {
		fmt.Printf("As integer: %d\nDoubled: %d\nSquared: %d\n",
			numInt, numInt*2, numInt*numInt)
	}

	// 2. Advanced Number Systems
	fmt.Println("\n\033[1;32m2. NUMBER SYSTEM CONVERSIONS\033[0m")
	fmt.Println("---------------------------")

	fmt.Print("Enter a decimal number for conversion: ")
	input, _ = reader.ReadString('\n')
	if num, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64); err == nil {
		// Convert to different number systems
		fmt.Printf("Binary (base 2):  %s\n", strconv.FormatInt(num, 2))
		fmt.Printf("Octal (base 8):   %s\n", strconv.FormatInt(num, 8))
		fmt.Printf("Hexadecimal:      %s\n", strconv.FormatInt(num, 16))
		fmt.Printf("Base-36:          %s\n", strconv.FormatInt(num, 36))
	}

	// 3. String Manipulations and Conversions
	fmt.Println("\n\033[1;32m3. STRING MANIPULATIONS\033[0m")
	fmt.Println("---------------------------")

	fmt.Print("Enter a sentence: ")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	// String analysis
	fmt.Printf("Original:      %s\n", sentence)
	fmt.Printf("Uppercase:     %s\n", strings.ToUpper(sentence))
	fmt.Printf("Lowercase:     %s\n", strings.ToLower(sentence))
	fmt.Printf("Word Count:    %d\n", len(strings.Fields(sentence)))
	fmt.Printf("Character Count: %d\n", len(sentence))

	// 4. Currency and Price Calculations
	fmt.Println("\n\033[1;32m4. CURRENCY CALCULATIONS\033[0m")
	fmt.Println("---------------------------")

	fmt.Print("Enter price in dollars (e.g., 49.99): ")
	input, _ = reader.ReadString('\n')
	if price, err := strconv.ParseFloat(strings.TrimSpace(input), 64); err == nil {
		// Calculate various amounts
		tax := price * 0.20      // 20% tax
		discount := price * 0.15 // 15% discount
		final := price + tax - discount

		fmt.Printf("\nPrice Breakdown:\n")
		fmt.Printf("Original Price: $%8.2f\n", price)
		fmt.Printf("Tax (20%%):     $%8.2f\n", tax)
		fmt.Printf("Discount (15%%): $%8.2f\n", discount)
		fmt.Printf("Final Price:    $%8.2f\n", final)

		// Convert to cents (integer)
		cents := int64(final * 100)
		fmt.Printf("In cents:       %d¢\n", cents)
	}

	// 5. Time and Date Conversions
	fmt.Println("\n\033[1;32m5. DATE AND TIME CONVERSIONS\033[0m")
	fmt.Println("---------------------------")

	now := time.Now()
	fmt.Printf("Current Time:     %v\n", now)
	fmt.Printf("Unix Timestamp:   %d\n", now.Unix())
	fmt.Printf("Formatted Date:   %s\n", now.Format("2006-01-02"))
	fmt.Printf("Formatted Time:   %s\n", now.Format("15:04:05"))
	fmt.Printf("Custom Format:    %s\n", now.Format("Monday, January 2, 2006"))

	// 6. Boolean Conversions and Logic
	fmt.Println("\n\033[1;32m6. BOOLEAN OPERATIONS\033[0m")
	fmt.Println("---------------------------")

	fmt.Print("Enter a boolean value (true/false): ")
	input, _ = reader.ReadString('\n')
	if boolVal, err := strconv.ParseBool(strings.TrimSpace(input)); err == nil {
		fmt.Printf("Boolean value:   %v\n", boolVal)
		fmt.Printf("Negated value:   %v\n", !boolVal)
		fmt.Printf("As integer:      %d\n", map[bool]int{false: 0, true: 1}[boolVal])
		fmt.Printf("As string:       %s\n", strconv.FormatBool(boolVal))
	}

	// 7. Multi-value Processing
	fmt.Println("\n\033[1;32m7. MULTI-VALUE PROCESSING\033[0m")
	fmt.Println("---------------------------")

	fmt.Println("Enter three numbers separated by spaces (e.g., 10 20.5 30):")
	input, _ = reader.ReadString('\n')
	numbers := strings.Fields(strings.TrimSpace(input))

	if len(numbers) == 3 {
		// Try converting each number
		if num1, err1 := strconv.Atoi(numbers[0]); err1 == nil {
			if num2, err2 := strconv.ParseFloat(numbers[1], 64); err2 == nil {
				if num3, err3 := strconv.Atoi(numbers[2]); err3 == nil {
					// Perform calculations
					sum := float64(num1) + num2 + float64(num3)
					avg := sum / 3.0

					fmt.Printf("\nResults:\n")
					fmt.Printf("Sum:     %.2f\n", sum)
					fmt.Printf("Average: %.2f\n", avg)
					fmt.Printf("Product: %.2f\n", float64(num1)*num2*float64(num3))
				}
			}
		}
	}

	// 8. Temperature Conversions
	fmt.Println("\n\033[1;32m8. TEMPERATURE CONVERSIONS\033[0m")
	fmt.Println("---------------------------")

	fmt.Print("Enter temperature in Celsius: ")
	input, _ = reader.ReadString('\n')
	if celsius, err := strconv.ParseFloat(strings.TrimSpace(input), 64); err == nil {
		// Convert to other scales
		fahrenheit := (celsius * 9 / 5) + 32
		kelvin := celsius + 273.15

		fmt.Printf("\nTemperature Conversions:\n")
		fmt.Printf("Celsius:    %.2f°C\n", celsius)
		fmt.Printf("Fahrenheit: %.2f°F\n", fahrenheit)
		fmt.Printf("Kelvin:     %.2fK\n", kelvin)
	}

	// Program completion
	fmt.Println("\n\033[1;36m=== Program Complete ===\033[0m")
	fmt.Println("\033[1;33mPress Enter to exit...\033[0m")
	reader.ReadString('\n')
}
