package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{
		"saturday",
		"sunday",
		"monday",
		"tuesday",
		"wednesday",
		"thursday",
		"friday",
	}

	fmt.Println("\n1. Original array:")
	fmt.Println(days)

	fmt.Println("\n2. Basic for loop with comma:")
	for d := 0; d < len(days); d++ {
		fmt.Print(days[d], ",")
	}
	fmt.Println()

	fmt.Println("\n3. For range with index:")
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("\n4. For loop without last comma:")
	for d := 0; d < len(days); d++ {
		if d == len(days)-1 {
			fmt.Print(days[d])
		} else {
			fmt.Print(days[d], ",")
		}
	}
	fmt.Println()

	fmt.Println("\n5. Using strings.Join:")
	fmt.Println(strings.Join(days, ","))

	fmt.Println("\n6. For range with index and value:")
	for i, day := range days {
		fmt.Printf("Index: %d, Value: %s\n", i, day)
	}

	fmt.Println("\n7. For range with only value:")
	for _, day := range days {
		fmt.Println(day)
	}

	fmt.Println("\n8. While style loop:")
	i := 0
	for i < len(days) {
		fmt.Println(days[i])
		i++
	}

	fmt.Println("\n9. Infinite loop with break:")
	j := 0
	for {
		if j >= len(days) {
			break
		}
		fmt.Println(days[j])
		j++
	}

	fmt.Println("\n10. Loop with continue (skipping sunday):")
	for _, day := range days {
		if day == "sunday" {
			continue
		}
		fmt.Println(day)
	}

	fmt.Println("\n11. Reverse loop:")
	for i := len(days) - 1; i >= 0; i-- {
		fmt.Println(days[i])
	}

	fmt.Println("\n12. Loop with custom step size (print every second day):")
	for i := 0; i < len(days); i += 2 {
		fmt.Println(days[i])
	}
}
