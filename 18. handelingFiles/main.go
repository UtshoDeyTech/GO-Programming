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
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

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
	checkNilErr(err)

	fmt.Println("The text data inside the file is-\n", string(databyte))
}

func appendToFile(filename, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	checkNilErr(err)
	defer file.Close()

	_, err = file.WriteString(content)
	checkNilErr(err)
	fmt.Println("Successfully appended data to file")
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
