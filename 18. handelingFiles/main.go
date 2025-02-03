package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files handling in golang")

	content := "This needs to go in a file - utshodey.tech"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("File length is:", length)
	defer file.Close()

	readFile("./myfile.txt")
	fmt.Println()

	// Demonstrating append to file
	appendToFile("./myfile.txt", "\nAppending some extra content!")
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("The text data inside the file is-\n", string(databyte))
}

func appendToFile(filename, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully appended data to file")
}
