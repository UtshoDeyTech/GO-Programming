package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma or || err
	// like try and catch

	input, _ := reader.ReadString('\n')
	fmt.Println("Thank for rating, ", input)
	fmt.Printf("Type of the rating is %T ", input)

	_, err := reader.ReadString('\n')
	fmt.Println("Thank for rating, ", err)
	fmt.Printf("Type of the rating is %T ", err)
}
