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

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println("Error Decoding", err)
		return
	}
	newUser.ID = idCounter
	idCounter++
	users = append(users, newUser)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(newUser)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idstr := r.URL.Query().Get("id")
	if idstr != "" {
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
	json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("Error ID:", err)
		return
	}

	var updatedUser User
	err1 := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err1 != nil {
		fmt.Println("Error decoding:", err)
		return
	}
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

func UserManagement() {

	http.HandleFunc("/users/", userHandler)
	fmt.Println("Server running on Path 8080.")
	http.ListenAndServe(":8080", nil)
}
