package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON in golang")
	EncodeJson()

	DecodeJson()
}

func EncodeJson() {
	affpilotUsers := []course{
		{"utsho dey", 1200, "affpilot.com", "123456", []string{"js", "python"}},
		{"sanjib sen", 1400, "affpilot.com", "123456", []string{"js", "go"}},
		{"kabir", 1500, "affpilot.com", "123456", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(affpilotUsers, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(finalJson))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "utsho dey",
                "price": 1200,
                "website": "affpilot.com",
                "tags": ["js","python"]
        }
	`)

	var affpilotCourse course

	checkvalid := json.Valid(jsonDataFromWeb)

	if checkvalid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &affpilotCourse)
		fmt.Printf("%#v\n", affpilotCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// somecases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", affpilotCourse)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v \n Type is: %T", k, v, v)
	}
}
