package main

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
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

// Global variables
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", CreatedAt: time.Now().Format(time.RFC3339)},
}

// Logger function
func logEndpoint(r *http.Request, startTime time.Time, statusCode int) {
	duration := time.Since(startTime)

	fmt.Printf("\n=== Request Log ===\n")
	fmt.Printf("Timestamp: %v\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("Endpoint: %s\n", r.URL.Path)
	fmt.Printf("Status: %d\n", statusCode)
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Client IP: %s\n", r.RemoteAddr)
	fmt.Printf("================\n")
}

// GET handler function
func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		logEndpoint(r, startTime, http.StatusMethodNotAllowed)
		return
	}

	response := Response{
		Message: "Users retrieved successfully",
		Status:  http.StatusOK,
		Data:    users,
	}

	sendJSONResponse(w, response)
	logEndpoint(r, startTime, http.StatusOK)
}

// POST handler function
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		logEndpoint(r, startTime, http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logEndpoint(r, startTime, http.StatusBadRequest)
		return
	}

	newUser.ID = getNextUserID()
	newUser.CreatedAt = time.Now().Format(time.RFC3339)
	users = append(users, newUser)

	saveUsersToFile()

	response := Response{
		Message: "User created successfully",
		Status:  http.StatusCreated,
		Data:    newUser,
	}

	w.WriteHeader(http.StatusCreated)
	sendJSONResponse(w, response)
	logEndpoint(r, startTime, http.StatusCreated)
}

// PUT handler function
func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		logEndpoint(r, startTime, http.StatusMethodNotAllowed)
		return
	}

	id := getUserIDFromURL(r.URL.Path, "/api/put/")
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logEndpoint(r, startTime, http.StatusBadRequest)
		return
	}

	if !updateUserByID(id, &updatedUser) {
		http.Error(w, "User not found", http.StatusNotFound)
		logEndpoint(r, startTime, http.StatusNotFound)
		return
	}

	saveUsersToFile()
	sendJSONResponse(w, Response{
		Message: "User updated successfully",
		Status:  http.StatusOK,
		Data:    updatedUser,
	})
	logEndpoint(r, startTime, http.StatusOK)
}

// DELETE handler function
func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		logEndpoint(r, startTime, http.StatusMethodNotAllowed)
		return
	}

	id := getUserIDFromURL(r.URL.Path, "/api/delete/")
	if !deleteUserByID(id) {
		http.Error(w, "User not found", http.StatusNotFound)
		logEndpoint(r, startTime, http.StatusNotFound)
		return
	}

	saveUsersToFile()
	sendJSONResponse(w, Response{
		Message: "User deleted successfully",
		Status:  http.StatusOK,
	})
	logEndpoint(r, startTime, http.StatusOK)
}

// Form data handler function
func handleFormData(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		logEndpoint(r, startTime, http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logEndpoint(r, startTime, http.StatusBadRequest)
		return
	}

	newUser := createUserFromForm(r)
	users = append(users, newUser)
	saveUsersToFile()
	saveFormToFile(newUser)

	sendJSONResponse(w, Response{
		Message: "Form data processed successfully",
		Status:  http.StatusOK,
		Data:    newUser,
	})
	logEndpoint(r, startTime, http.StatusOK)
}

// File upload handler function
func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		logEndpoint(r, startTime, http.StatusMethodNotAllowed)
		return
	}

	file, handler, err := processFileUpload(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logEndpoint(r, startTime, http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := saveUploadedFile(file, handler); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logEndpoint(r, startTime, http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, Response{
		Message: fmt.Sprintf("File %s uploaded successfully", handler.Filename),
		Status:  http.StatusOK,
	})
	logEndpoint(r, startTime, http.StatusOK)
}

// [Rest of the helper functions remain the same]
func getNextUserID() int {
	maxID := 0
	for _, user := range users {
		if user.ID > maxID {
			maxID = user.ID
		}
	}
	return maxID + 1
}

func getUserIDFromURL(path, prefix string) string {
	return path[len(prefix):]
}

func updateUserByID(id string, updatedUser *User) bool {
	for i, user := range users {
		if fmt.Sprint(user.ID) == id {
			updatedUser.ID = user.ID
			updatedUser.CreatedAt = user.CreatedAt
			users[i] = *updatedUser
			return true
		}
	}
	return false
}

func deleteUserByID(id string) bool {
	for i, user := range users {
		if fmt.Sprint(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}

func createUserFromForm(r *http.Request) User {
	return User{
		ID:        getNextUserID(),
		Name:      r.FormValue("name"),
		Email:     r.FormValue("email"),
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}

func processFileUpload(r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return nil, nil, err
	}
	return r.FormFile("file")
}

func saveUploadedFile(file multipart.File, handler *multipart.FileHeader) error {
	os.MkdirAll("uploads", os.ModePerm)
	dst, err := os.Create(fmt.Sprintf("uploads/%s", handler.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}

func saveUsersToFile() {
	data, _ := json.MarshalIndent(users, "", "    ")
	os.WriteFile("users.json", data, 0644)
}

func saveFormToFile(user User) {
	formData := fmt.Sprintf("ID: %d\nName: %s\nEmail: %s\nCreated At: %s\n\n",
		user.ID, user.Name, user.Email, user.CreatedAt)
	os.WriteFile("form_submissions.txt", []byte(formData), 0644)
}

func sendJSONResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func setupRoutes() {
	fmt.Println("Setting up routes...")
	http.HandleFunc("/api/get", handleGetUsers)
	http.HandleFunc("/api/post", handleCreateUser)
	http.HandleFunc("/api/put/", handleUpdateUser)
	http.HandleFunc("/api/delete/", handleDeleteUser)
	http.HandleFunc("/api/form", handleFormData)
	http.HandleFunc("/api/upload", handleFileUpload)
}

func main() {
	setupRoutes()
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
