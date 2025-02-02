package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string) // [key_type]value_type

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS shorts for:", languages["js"])

	delete(languages, "rb")
	fmt.Println("List of all languages:", languages)

	// loops are insteresting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
