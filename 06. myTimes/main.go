package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2025 12-12-12 Monday"))

	createDate := time.Date(2020, time.June, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2025 Monday"))
}
