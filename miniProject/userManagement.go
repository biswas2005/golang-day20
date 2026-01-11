package miniProject

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "abc", Email: "abc@gmail.com"},
}
var idCounter = 2

// createUser() creates a new User
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, `{"Error":"Invalid Body"}`, http.StatusBadRequest)
		return
	}
	//validation is checked
	if !validation(newUser, w) {
		return
	}
	newUser.ID = idCounter
	idCounter++
	//newUser is added to the slice
	users = append(users, newUser)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(newUser)
}

// getUser() reads the users from the Slice
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idstr := r.URL.Query().Get("id")
	if idstr != "" {
		//convert idstr from string to int
		id, err := strconv.Atoi(idstr)
		if err != nil {
			fmt.Println("Error converting ID:", err)
			return
		}

		for _, user := range users {
			if user.ID == id {
				json.NewEncoder(w).Encode(user)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//reads the users
	json.NewEncoder(w).Encode(users)
}

// updateUser() changes the value
// it takes id as an identifier
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, `{"Error":"Invalid ID"}`, http.StatusBadRequest)
		return
	}
	//updatedUser has the new data
	var updatedUser User
	err1 := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err1 != nil {
		http.Error(w, `{"Error":"Invalid Body"}`, http.StatusBadRequest)
		return
	}
	//checks for the validation
	if !validation(updatedUser, w) {
		return
	}
	//updates the data
	for i, user := range users {
		if user.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// deleteUser() deletes if data exist
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("Error updating:", err)
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"message":"user %d deleted}`, id)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

// userHandler() switches the operation
func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		createUser(w, r)

	case http.MethodGet:
		getUser(w, r)

	case http.MethodPut:
		updateUser(w, r)

	case http.MethodDelete:
		deleteUser(w, r)

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// validation() checks for some cases
func validation(user User, w http.ResponseWriter) bool {
	w.Header().Set("Content-Type", "application/json")
	//name cannot be empty
	if strings.TrimSpace(user.Name) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Error": "Name cannot be empty."})
		return false
	}
	//email cannot be empty
	if strings.TrimSpace(user.Email) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Error": "Email cannot be empty."})
		return false
	}
	//suffix has to @gmail.com for email
	if !strings.HasSuffix(user.Email, "@gmail.com") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Error": "Wrong Email format."})
		return false
	}
	//prefix must exist
	prefix := strings.TrimSuffix(user.Email, "@gmail.com")
	if prefix == "" {
		http.Error(w, `{"Error:mail cannot be empty"}`, http.StatusBadRequest)
		return false
	}
	return true
}

func UserManagement() {
	//Acts as a signal
	//control the execution
	http.HandleFunc("/users/", userHandler)
	fmt.Println("Server running on Path 8080.")
	http.ListenAndServe(":8080", nil)
}
