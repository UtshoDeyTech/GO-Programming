package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://simpletext.app/v5/"

func main() {
	fmt.Println("Web request handle in golang")

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	content := string(databyte)
	fmt.Println(content)

	// Writing response to a file
	saveToFile("response.txt", content)
	fmt.Println("Response saved to response.txt")

	// Reading the saved file
	readFile("response.txt")
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func saveToFile(filename, data string) {
	file, err := os.Create(filename)
	checkNilErr(err)
	defer file.Close()

	_, err = io.WriteString(file, data)
	checkNilErr(err)
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("File content is:\n", string(databyte))
}
