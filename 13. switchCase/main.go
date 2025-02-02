package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	// Multiple dice rolls
	for i := 0; i < 3; i++ {
		fmt.Printf("\nRoll #%d:\n", i+1)
		rollDice()
	}

	// Multiple dice at once
	fmt.Println("\nRolling two dice at once:")
	rollMultipleDice(2)

	// Switch with multiple cases
	fmt.Println("\nRolling with multiple case matching:")
	rollDiceWithMultipleCases()

	// Switch with ranges
	fmt.Println("\nRolling with range checks:")
	rollDiceWithRange()

	// Switch with interface{}
	fmt.Println("\nDifferent types of rolls:")
	differentTypesOfRolls()
}

func rollDice() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 and move 2 spots")
	case 3:
		fmt.Println("Dice value is 3 and move 3 spots")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 and move 4 spots")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 and move 5 spots")
	case 6:
		fmt.Println("Dice value is 6 and you get another chance to roll")
		rollDice() // Recursive call for rolling again
	}
}

func rollMultipleDice(numDice int) {
	rand.Seed(time.Now().UnixNano())
	total := 0

	for i := 0; i < numDice; i++ {
		value := rand.Intn(6) + 1
		total += value
		fmt.Printf("Dice %d: %d\n", i+1, value)
	}

	fmt.Printf("Total value: %d\n", total)
}

func rollDiceWithMultipleCases() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1, 2:
		fmt.Println("Bad luck! You got a low number")
	case 3, 4:
		fmt.Println("Not bad! Medium number")
	case 5, 6:
		fmt.Println("Great! High number")
	}
}

func rollDiceWithRange() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch {
	case diceNumber <= 2:
		fmt.Println("Low roll")
	case diceNumber <= 4:
		fmt.Println("Medium roll")
	default:
		fmt.Println("High roll")
	}
}

func differentTypesOfRolls() {
	rand.Seed(time.Now().UnixNano())

	// Different types of values
	values := []interface{}{
		rand.Intn(6) + 1,
		"special",
		true,
		3.14,
	}

	for _, value := range values {
		fmt.Printf("\nRolling with value: %v\n", value)
		switch v := value.(type) {
		case int:
			fmt.Printf("Numeric roll: %d\n", v)
		case string:
			fmt.Printf("Special roll: %s\n", v)
		case bool:
			fmt.Printf("Boolean roll: %v\n", v)
		default:
			fmt.Printf("Unknown type roll: %v\n", v)
		}
	}
}
