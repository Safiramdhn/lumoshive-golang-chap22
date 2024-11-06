package handlers

import (
	"golang-beginner-22/database"
	"golang-beginner-22/models"
	"golang-beginner-22/repositories"
	"golang-beginner-22/services"
	"golang-beginner-22/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)

		return
	}

	if err := r.ParseForm(); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	description := r.FormValue("description")
	todoInput := models.Todos{
		Description: description,
	}

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
	_, err = todoService.CreateTodo(&todoInput, token)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	http.Redirect(w, r, "/todo-list", http.StatusFound)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)

		return
	}

	if err := r.ParseForm(); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	todoID := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(todoID)
	// Parse the form data (status value)
	r.ParseForm()
	newStatus := r.FormValue("status")

	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	todoRepo := repositories.NewTodoRepositoryDB(db)
	todoService := services.NewTodoService(*todoRepo)
	err = todoService.UpdateTodo(id, newStatus)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	http.Redirect(w, r, "/todo-list", http.StatusFound)
}
