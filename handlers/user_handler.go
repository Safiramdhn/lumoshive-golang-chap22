package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-beginner-22/database"
	"golang-beginner-22/models"
	"golang-beginner-22/repositories"
	"golang-beginner-22/services"
)

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	// Check if the request method is POST
// 	if r.Method != http.MethodPost {
// 		return
// 	}

// 	// Decode the request body
// 	var login models.User
// 	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
// 		return
// 	}

// 	// Initialize the database
// 	db, err := database.InitDB()
// 	if err != nil {

// 		return
// 	}
// 	defer db.Close()

// 	// Create repository and service
// 	userRepo := repositories.NewUserRepositoryDB(db)
// 	userService := services.NewUserService(*userRepo)

// 	Authenticate the user
// 	user, err = userService.LoginService(login.Email, login.Password)
// 	// if err != nil {

// 	// 	return
// 	// }
// }

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		fmt.Printf("%d : Invalid request method\n", http.StatusMethodNotAllowed)
		return
	}

	var userInput models.User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		fmt.Printf("%d : Invalid request payload\n", http.StatusBadRequest)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		fmt.Printf("%d : Database connection error: %v\n", http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	userRepo := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(*userRepo)
	newUser, err := userService.CreateUser(&userInput)
	if err != nil {
		fmt.Printf("%d : error creating user: %v\n", http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("token", newUser.Token)
	fmt.Printf("%d : User create successfully\n", http.StatusOK)
}
