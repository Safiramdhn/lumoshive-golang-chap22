package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	w.Header().Add("token", newUser.Token)
	fmt.Printf("%d : User create successfully\n", http.StatusOK)
	http.Redirect(w, r, "/todo-list", http.StatusOK)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB()
	if err != nil {
		fmt.Printf("%d : Database connection error: %v\n", http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	userRepo := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(*userRepo)
	users, err := userService.GetAllUsers()
	if err != nil {
		fmt.Printf("%d : error fetching users: %v\n", http.StatusInternalServerError, err.Error())
		return
	}
	RenderTemplate(w, "list_data_user.html", users)
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id_int := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(id_int)

	db, err := database.InitDB()
	if err != nil {
		fmt.Printf("%d : Database connection error: %v\n", http.StatusInternalServerError, err.Error())
		return
	}

	userRepo := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(*userRepo)
	user, err := userService.GetUserByID(id)
	if err != nil {
		fmt.Printf("%d : error fetching user: %v\n", http.StatusInternalServerError, err.Error())
		return
	}
	RenderTemplate(w, "user_details.html", user)
}
