package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"golang-beginner-22/database"
	// "golang-beginner-22/models"
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
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Populate user input data
	// userInput := models.User{
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	// }
	// fmt.Printf("input: %v\n", userInput)

	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Create user repository and service
	userRepo := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(*userRepo)

	// Attempt to create a new user
	newUser, err := userService.CreateUser(name, email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the token in the response header
	w.Header().Set("token", newUser.Token)

	// Redirect to the todo list with a success message
	http.Redirect(w, r, "/todo/todo-list", http.StatusSeeOther)
	fmt.Printf("User created successfully, status: %d\n", http.StatusOK)
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
