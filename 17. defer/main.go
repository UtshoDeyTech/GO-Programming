package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")
	myDefer()
	fmt.Println()
	simpleDeferExample()
	fmt.Println()
	fileOperationSimulation()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, " ")
	}
}

// Simple defer example
func simpleDeferExample() {
	fmt.Println("Start simpleDeferExample")
	defer fmt.Println("This will be printed last in simpleDeferExample")
	fmt.Println("End simpleDeferExample")
}

// Simulating file operation with defer
func fileOperationSimulation() {
	fmt.Println("Opening file")
	defer fmt.Println("Closing file")
	fmt.Println("Reading file")
	fmt.Println("Processing file")
}
