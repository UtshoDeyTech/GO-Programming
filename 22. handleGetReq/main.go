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

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func main() {
	for {
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
		case "7":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}

// 1. GET - Get all users
func getAllUsers() {
	resp, err := http.Get("http://localhost:8080/api/get")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var users []User
	json.NewDecoder(resp.Body).Decode(&users)

	fmt.Println("\nUsers List:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Created: %s\n",
			user.ID, user.Name, user.Email, user.CreatedAt)
	}
}

// 2. POST - Create new user
func createUser() {
	var name, email string
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter email: ")
	fmt.Scanln(&email)

	newUser := User{Name: name, Email: email}
	jsonData, _ := json.Marshal(newUser)

	resp, err := http.Post("http://localhost:8080/api/post",
		"application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var createdUser User
	json.NewDecoder(resp.Body).Decode(&createdUser)
	fmt.Printf("\nCreated User - ID: %d, Name: %s, Email: %s\n",
		createdUser.ID, createdUser.Name, createdUser.Email)
}

// 3. PUT - Update user
func updateUser() {
	var id, name, email string
	fmt.Print("Enter user ID to update: ")
	fmt.Scanln(&id)
	fmt.Print("Enter new name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter new email: ")
	fmt.Scanln(&email)

	updatedUser := User{Name: name, Email: email}
	jsonData, _ := json.Marshal(updatedUser)

	req, _ := http.NewRequest(http.MethodPut,
		"http://localhost:8080/api/put/"+id,
		bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User updated successfully!")
	} else {
		fmt.Println("Error updating user!")
	}
}

// 4. DELETE - Delete user
func deleteUser() {
	var id string
	fmt.Print("Enter user ID to delete: ")
	fmt.Scanln(&id)

	req, _ := http.NewRequest(http.MethodDelete,
		"http://localhost:8080/api/delete/"+id,
		nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User deleted successfully!")
	} else {
		fmt.Println("Error deleting user!")
	}
}

// 5. FORM - Submit form data
func submitForm() {
	var name, email string
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter email: ")
	fmt.Scanln(&email)

	formData := url.Values{
		"name":  {name},
		"email": {email},
	}

	resp, err := http.PostForm("http://localhost:8080/api/form", formData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var result User
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Printf("\nForm Submitted - ID: %d, Name: %s, Email: %s\n",
		result.ID, result.Name, result.Email)
}

// 6. UPLOAD - Upload file
func uploadFile() {
	fmt.Print("Enter file path to upload: ")
	var filePath string
	fmt.Scanln(&filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filePath[strings.LastIndex(filePath, "/")+1:])
	if err != nil {
		fmt.Printf("Error creating form file: %v\n", err)
		return
	}

	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", "http://localhost:8080/api/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("File uploaded successfully!")
	} else {
		fmt.Println("Error uploading file!")
	}
}
