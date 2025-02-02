package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list length is:", len(fruitList))

	var vegList = [3]string{"Potato", "beans", "mushroom"}
	fmt.Println("Vagitable list is:", vegList)
	fmt.Println("Vagitable list length is:", len(vegList))
}
