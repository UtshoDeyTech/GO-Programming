package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// User represents our data structure
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// Simple in-memory database
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", CreatedAt: time.Now().Format(time.RFC3339)},
}

func main() {
	// 1. GET endpoint - Get all users
	http.HandleFunc("/api/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	// 2. POST endpoint - Create new user
	http.HandleFunc("/api/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Set user details
		newUser.ID = len(users) + 1
		newUser.CreatedAt = time.Now().Format(time.RFC3339)
		users = append(users, newUser)

		// Save to JSON file
		saveUsersToFile()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	})

	// 3. PUT endpoint - Update user
	http.HandleFunc("/api/put/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Get user ID from URL
		id := r.URL.Path[len("/api/put/"):]
		var updatedUser User
		if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update user
		found := false
		for i := range users {
			if fmt.Sprint(users[i].ID) == id {
				updatedUser.ID = users[i].ID
				updatedUser.CreatedAt = users[i].CreatedAt
				users[i] = updatedUser
				found = true
				break
			}
		}

		if !found {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		saveUsersToFile()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedUser)
	})

	// 4. DELETE endpoint - Delete user
	http.HandleFunc("/api/delete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Path[len("/api/delete/"):]
		found := false
		for i := range users {
			if fmt.Sprint(users[i].ID) == id {
				users = append(users[:i], users[i+1:]...)
				found = true
				break
			}
		}

		if !found {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		saveUsersToFile()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
	})

	// 5. FormData endpoint - Handle form submissions
	http.HandleFunc("/api/form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get form data
		newUser := User{
			ID:        len(users) + 1,
			Name:      r.FormValue("name"),
			Email:     r.FormValue("email"),
			CreatedAt: time.Now().Format(time.RFC3339),
		}

		users = append(users, newUser)
		saveUsersToFile()

		// Save form data to a text file
		formData := fmt.Sprintf("ID: %d\nName: %s\nEmail: %s\nCreated At: %s\n\n",
			newUser.ID, newUser.Name, newUser.Email, newUser.CreatedAt)
		os.WriteFile("form_submissions.txt", []byte(formData), 0644)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newUser)
	})

	// 6. File Upload endpoint
	http.HandleFunc("/api/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse multipart form with 10 MB max memory
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get file from form
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create uploads directory if it doesn't exist
		os.MkdirAll("uploads", os.ModePerm)

		// Create new file in uploads directory
		dst, err := os.Create(fmt.Sprintf("uploads/%s", handler.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy uploaded file to destination
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("File %s uploaded successfully", handler.Filename),
		})
	})

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// Helper function to save users to a JSON file
func saveUsersToFile() {
	data, _ := json.MarshalIndent(users, "", "    ")
	os.WriteFile("users.json", data, 0644)
}
