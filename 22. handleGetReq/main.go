package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Data Structures
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    any    `json:"data,omitempty"`
}

const baseURL = "http://localhost:8080/api"

func main() {
	for {
		choice := displayMenu()
		if choice == "7" {
			fmt.Println("Exiting...")
			break
		}
		handleUserChoice(choice)
	}
}

// Menu Functions
func displayMenu() string {
	fmt.Println("\n=== API Testing Menu ===")
	fmt.Println("1. GET - Get all users")
	fmt.Println("2. POST - Create new user")
	fmt.Println("3. PUT - Update user")
	fmt.Println("4. DELETE - Delete user")
	fmt.Println("5. FORM - Submit form data")
	fmt.Println("6. UPLOAD - Upload file")
	fmt.Println("7. Exit")
	fmt.Print("Choose an option (1-7): ")

	var choice string
	fmt.Scanln(&choice)
	return choice
}

func handleUserChoice(choice string) {
	switch choice {
	case "1":
		getAllUsers()
	case "2":
		createUser()
	case "3":
		updateUser()
	case "4":
		deleteUser()
	case "5":
		submitForm()
	case "6":
		uploadFile()
	default:
		fmt.Println("Invalid option!")
	}
}

// API Functions
func getAllUsers() {
	fmt.Println("\nGetting all users...")
	resp, err := makeGetRequest("/get")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Printf("Error parsing response: %v\n", err)
		return
	}

	displayUsers(response.Data)
}

func createUser() {
	fmt.Println("\nCreating new user...")
	userData := getUserInput()
	jsonData, err := json.Marshal(userData)
	if err != nil {
		fmt.Printf("Error creating JSON: %v\n", err)
		return
	}

	resp, err := makePostRequest("/post", jsonData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	displayResponse(resp)
}

func updateUser() {
	fmt.Println("\nUpdating user...")
	id := getInputString("Enter user ID to update: ")
	userData := getUserInput()
	jsonData, err := json.Marshal(userData)
	if err != nil {
		fmt.Printf("Error creating JSON: %v\n", err)
		return
	}

	resp, err := makePutRequest("/put/"+id, jsonData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	displayResponse(resp)
}

func deleteUser() {
	fmt.Println("\nDeleting user...")
	id := getInputString("Enter user ID to delete: ")

	resp, err := makeDeleteRequest("/delete/" + id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	displayResponse(resp)
}

func submitForm() {
	fmt.Println("\nSubmitting form data...")
	formData := url.Values{
		"name":  {getInputString("Enter name: ")},
		"email": {getInputString("Enter email: ")},
	}

	resp, err := http.PostForm(baseURL+"/form", formData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	displayResponse(resp)
}

func uploadFile() {
	fmt.Println("\nUploading file...")
	filePath := getInputString("Enter file path to upload: ")
	if err := handleFileUpload(filePath); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// Helper Functions
func getUserInput() User {
	return User{
		Name:  getInputString("Enter name: "),
		Email: getInputString("Enter email: "),
	}
}

func getInputString(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

func handleFileUpload(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filePath[strings.LastIndex(filePath, "/")+1:])
	if err != nil {
		return fmt.Errorf("error creating form file: %v", err)
	}

	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", baseURL+"/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error uploading file: %v", err)
	}
	defer resp.Body.Close()

	displayResponse(resp)
	return nil
}

// Display Functions
func displayUsers(data interface{}) {
	users, ok := data.([]interface{})
	if !ok {
		fmt.Println("No users found")
		return
	}

	fmt.Println("\nUsers List:")
	for _, u := range users {
		user, ok := u.(map[string]interface{})
		if ok {
			fmt.Printf("ID: %v\nName: %v\nEmail: %v\nCreated At: %v\n\n",
				user["id"], user["name"], user["email"], user["created_at"])
		}
	}
}

func displayResponse(resp *http.Response) {
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Printf("Error parsing response: %v\n", err)
		return
	}

	fmt.Printf("\nResponse: %s (Status: %d)\n", response.Message, response.Status)
	if response.Data != nil {
		fmt.Printf("Data: %+v\n", response.Data)
	}
}

// HTTP Request Functions
func makeGetRequest(endpoint string) (*http.Response, error) {
	return http.Get(baseURL + endpoint)
}

func makePostRequest(endpoint string, jsonData []byte) (*http.Response, error) {
	return http.Post(baseURL+endpoint, "application/json", bytes.NewBuffer(jsonData))
}

func makePutRequest(endpoint string, jsonData []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, baseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func makeDeleteRequest(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	return client.Do(req)
}
