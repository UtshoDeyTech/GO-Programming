package main

import "fmt"

const LoginToken string = "sbfjkskds" // public
var jwtToken int = 30000              // public

func main() {
	var username string = "utsho"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float32 = 255.387478242492378
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	var longFloat float64 = 255.387478242492378
	fmt.Println(longFloat)
	fmt.Printf("Variable is of type : %T \n", longFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// implicit type
	var website = "utshodey.tech"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)

	// no var style
	numberofUser := 300000.0
	fmt.Println(numberofUser)
	fmt.Printf("Variable is of type : %T \n", website)
}
