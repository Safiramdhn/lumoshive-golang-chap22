package handlers

import (
	"golang-beginner-22/database"
	"golang-beginner-22/repositories"
	"golang-beginner-22/services"
	"golang-beginner-22/utils"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var templates = template.Must(template.ParseGlob("views/*.html"))

// RegisterFormHandler handles the registration page request
func RegisterFormHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "registration", map[string]interface{}{
		"Title": "Registration",
	})
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login", map[string]interface{}{
		"Title": "Login",
	})
}

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	userRepo := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(*userRepo)
	users, err := userService.GetAllUsers()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	templates.ExecuteTemplate(w, "user-list", map[string]interface{}{
		"Users": users,
	})
}

func UserTodoListHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil || cookie.Value == "" {
		utils.RespondWithJSON(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	token := cookie.Value
	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	todoRepo := repositories.NewTodoRepositoryDB(db)
	todoService := services.NewTodoService(*todoRepo)
	todoList, err := todoService.GetTodosByToken(token)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	templates.ExecuteTemplate(w, "todo-list", map[string]interface{}{
		"Todos": todoList,
	})
}

func UserDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)

		return
	}

	userId := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(userId)

	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	userRepo := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(*userRepo)
	user, err := userService.GetUserByID(id)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	templates.ExecuteTemplate(w, "user-detail", map[string]interface{}{
		"User": user,
	})
}
