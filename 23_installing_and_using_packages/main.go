package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserPreferences struct {
	PublicProfile bool `json:publicProfile`
	DarkMode      bool `json:darkMode`
}

type User struct {
	ID          int              `json:id`
	Username    string           `json:username`
	Email       string           `json:email`
	Preferences *UserPreferences `json:preferences` // like the settigs
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	// we need to set the return content to json, since we'll
	// be sending back json
	w.Header().Set("Content-Type", "application/json")

	// will respond with the json
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User

	// decode body and store the data in the user variable
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	user.ID = users[len(users)-1].ID + 1
	// any data that was sent in the body that is not part of the User
	// struct will be ignored and not added to the new user object
	users = append(users, user)

	json.NewEncoder(w).Encode((users))

	/* Example JSON body
	{
		"username": "Jackson",
		"email":    "jack@gmail.com",
		"preferences": {
			"publicProfile": true,
			"darkMode":      false
		}
	}
	*/
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // get params passed into URL
	userId, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid ID provided", http.StatusBadRequest)
		return
	}

	for index, user := range users {
		if user.ID == userId {
			// remove selected user
			users = append(users[:index], users[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode((users))
}

func main() {
	// set initial users
	users = append(users,
		User{
			ID:       1,
			Username: "Peter Pan",
			Email:    "peterpan@gmail.com",
			Preferences: &UserPreferences{
				PublicProfile: true,
				DarkMode:      false,
			},
		},
		User{
			ID:       2,
			Username: "Mike",
			Email:    "mike@gmail.com",
			Preferences: &UserPreferences{
				PublicProfile: false,
				DarkMode:      true,
			},
		},
		User{
			ID:       3,
			Username: "cooluser3",
			Email:    "cool@user.com",
			Preferences: &UserPreferences{
				PublicProfile: true,
				DarkMode:      true,
			},
		},
	)

	r := mux.NewRouter() // we now use mux to handle routes

	// routes
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Println("Starting server on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
