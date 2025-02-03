package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Part 1: Parsing a URL
	fmt.Println("=== URL Parsing Example ===")
	myurl := "https://example.com/path?key=value&other=param&fragment=ksjcf"

	// Parse the URL
	result, err := url.Parse(myurl)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	// Print URL components
	fmt.Println("URL Components:")
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("Raw Query:", result.RawQuery)

	// Parse and print query parameters
	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	// Print individual query parameters
	for key, values := range qparams {
		for _, value := range values {
			fmt.Printf("Key: %s, Value: %s\n", key, value)
		}
	}

	// Part 2: Constructing a URL
	fmt.Println("\n=== URL Construction Example ===")
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println("Constructed URL:", anotherURL)
}
