package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")

	reader := bufio.NewReader(os.Stdin)

	// Float conversion (your original code)
	fmt.Println("\n=== Rating Conversion ===")
	fmt.Println("Please rate our pizza between 1 to 5:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank for rating:", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error converting rating:", err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}

	// String to Integer conversion
	fmt.Println("\n=== Age Conversion ===")
	fmt.Println("Enter your age:")
	ageInput, _ := reader.ReadString('\n')
	age, err := strconv.Atoi(strings.TrimSpace(ageInput)) // Atoi = ASCII to Integer
	if err != nil {
		fmt.Println("Error converting age:", err)
	} else {
		fmt.Printf("Next year you will be: %d\n", age+1)
	}

	// String to Boolean conversion
	fmt.Println("\n=== Boolean Conversion ===")
	fmt.Println("Are you a vegetarian? (true/false):")
	boolInput, _ := reader.ReadString('\n')
	isVegetarian, err := strconv.ParseBool(strings.TrimSpace(boolInput))
	if err != nil {
		fmt.Println("Error converting boolean:", err)
	} else {
		if isVegetarian {
			fmt.Println("We'll suggest vegetarian pizzas for you!")
		} else {
			fmt.Println("You can try our meat pizzas!")
		}
	}

	// Integer to String conversion
	fmt.Println("\n=== Price Conversion ===")
	price := 1500
	priceStr := strconv.Itoa(price) // Integer to ASCII
	fmt.Printf("Price as string: %s, Type: %T\n", priceStr, priceStr)

	// Float to String with precision
	temperature := 98.6
	tempStr := fmt.Sprintf("%.1f", temperature)
	fmt.Printf("Temperature as string: %s, Type: %T\n", tempStr, tempStr)

	// String to Int64
	fmt.Println("\n=== Phone Number Conversion ===")
	fmt.Println("Enter your phone number (numbers only):")
	phoneInput, _ := reader.ReadString('\n')
	phoneNum, err := strconv.ParseInt(strings.TrimSpace(phoneInput), 10, 64)
	if err != nil {
		fmt.Println("Error converting phone number:", err)
	} else {
		fmt.Printf("Phone number as int64: %d\n", phoneNum)
	}

	// Base conversion examples
	fmt.Println("\n=== Base Conversions ===")
	// Decimal to Binary
	num := 42
	binary := strconv.FormatInt(int64(num), 2)
	fmt.Printf("%d in binary: %s\n", num, binary)

	// Decimal to Hexadecimal
	hex := strconv.FormatInt(int64(num), 16)
	fmt.Printf("%d in hexadecimal: %s\n", num, hex)

	// Binary to Decimal
	fmt.Println("Enter a binary number:")
	binInput, _ := reader.ReadString('\n')
	decimal, err := strconv.ParseInt(strings.TrimSpace(binInput), 2, 64)
	if err != nil {
		fmt.Println("Error converting binary:", err)
	} else {
		fmt.Printf("Binary to decimal: %d\n", decimal)
	}

	// String manipulation with conversion
	fmt.Println("\n=== Price Calculator ===")
	fmt.Println("Enter price (with decimal point):")
	priceInput, _ := reader.ReadString('\n')
	priceFloat, err := strconv.ParseFloat(strings.TrimSpace(priceInput), 64)
	if err != nil {
		fmt.Println("Error converting price:", err)
	} else {
		// Calculate tax (18%)
		tax := priceFloat * 0.18
		total := priceFloat + tax
		// Format currency with 2 decimal places
		fmt.Printf("Price: $%.2f\n", priceFloat)
		fmt.Printf("Tax (18%%): $%.2f\n", tax)
		fmt.Printf("Total: $%.2f\n", total)
	}

	// Multiple conversions in one operation
	fmt.Println("\n=== Order Processing ===")
	fmt.Println("Enter quantity and price separated by space (e.g., '2 10.5'):")
	orderInput, _ := reader.ReadString('\n')
	orderParts := strings.Fields(strings.TrimSpace(orderInput))

	if len(orderParts) == 2 {
		quantity, err1 := strconv.Atoi(orderParts[0])
		price, err2 := strconv.ParseFloat(orderParts[1], 64)

		if err1 == nil && err2 == nil {
			total := float64(quantity) * price
			fmt.Printf("Order total for %d items at $%.2f each: $%.2f\n",
				quantity, price, total)
		} else {
			fmt.Println("Error processing order input")
		}
	} else {
		fmt.Println("Invalid input format")
	}
}
