package handlers

import (
	"golang-beginner-22/database"
	"golang-beginner-22/models"
	"golang-beginner-22/repositories"
	"golang-beginner-22/services"
	"golang-beginner-22/utils"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)

		return
	}

	if err := r.ParseForm(); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	userInput := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	userRepo := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(*userRepo)
	newUser, err := userService.CreateUser(userInput)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, "User created successfully", newUser)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
		return
	}
	if err := r.ParseForm(); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	userRepo := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(*userRepo)
	user, err := userService.LoginService(email, password)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    user.Token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/user-list", http.StatusFound)
}
