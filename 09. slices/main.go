package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices on GO")

	var fruitList = []string{"Apple", "Mango", "Banana"}
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango", "Peach")
	fmt.Println(fruitList)

	fruitList = fruitList[1:]
	fmt.Println(fruitList)

	fruitList = fruitList[1:3]
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 954
	highScores[2] = 356
	highScores[3] = 845
	// highScores[4] = 545

	fmt.Println(highScores)

	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices baced on index
	var courses = []string{"reactjs", "django", "fastapi", "flask", "node", "express"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:2], courses[index+1:]...)
	fmt.Println(courses)
	fmt.Println(sort.StringsAreSorted(courses))

	sort.Strings(courses)
	fmt.Println(courses)
	fmt.Println(sort.StringsAreSorted(courses))
}
