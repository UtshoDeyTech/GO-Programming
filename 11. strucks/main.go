package main

import "fmt"

func main() {
	fmt.Println("Stucks in golnag")

	// no inheritance in golang; No super or parent
	user1 := User{"Utsho", "utshodey.tech@gmail.com", true, 26}
	fmt.Println(user1)
	fmt.Printf("user1's details are: %v\n", user1)
	fmt.Printf("user1's details are: %+v\n", user1)
	fmt.Printf("user1's Name is %v and Email is %v:", user1.Name, user1.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
