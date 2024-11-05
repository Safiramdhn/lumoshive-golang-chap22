package handlers

import (
	"golang-beginner-22/database"
	"golang-beginner-22/models"
	"golang-beginner-22/repositories"
	"golang-beginner-22/services"
	"log"
	"net/http"
	"strconv"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	var todoInput models.Todos
	todoInput.Description = r.PostForm.Get("description")
	token := r.Header.Get("token")

	db, err := database.InitDB()
	if err != nil {
		log.Printf("Database connection error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	todoRepo := repositories.NewTodoRepositoryDB(db)
	todoService := services.NewTodoService(*todoRepo)
	_, err = todoService.CreateTodo(&todoInput, token)
	if err != nil {
		log.Printf("Error creating todo: %v", err)
		http.Error(w, "Error creating todo", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/todo-list", http.StatusOK) // Reload the page
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		return
	}

	token := r.Header.Get("token")

	db, err := database.InitDB()
	if err != nil {

		return
	}

	todoRepo := repositories.NewTodoRepositoryDB(db)
	todoService := services.NewTodoService(*todoRepo)
	todos, err := todoService.GetTodosByUserId(token)
	if err != nil {
		return
	}
	RenderTemplate(w, "todo_list_html", todos)

}

// func GetTodoCountHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		return
// 	}

// 	token := r.Header.Get("token")
// 	db, err := database.InitDB()
// 	if err != nil {
// 		return
// 	}

// 	todoRepo := repositories.NewTodoRepositoryDB(db)
// 	todoService := services.NewTodoService(*todoRepo)
// 	todos, err := todoService.GetTodoCount(token)
// 	if err != nil {
// 		return
// 	}
// }

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todoInput models.Todos
	query := r.URL.Query()
	id_int, _ := strconv.Atoi(query.Get("id"))
	todoInput.ID = id_int
	todoInput.TodoStatus = query.Get("todo_status")

	db, err := database.InitDB()
	if err != nil {
		return
	}
	todoRepo := repositories.NewTodoRepositoryDB(db)
	todoService := services.NewTodoService(*todoRepo)
	_, err = todoService.UpdateTodo(&todoInput)
	if err != nil {
		return
	}
	http.Redirect(w, r, "/todo-list", http.StatusSeeOther) // Reload the page
}

// func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		return
// 	}

// 	var todo models.Todos
// 	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
// 		return
// 	}

// 	db, err := database.InitDB()
// 	if err != nil {
// 		return
// 	}
// 	todoRepo := repositories.NewTodoRepositoryDB(db)
// 	todoService := services.NewTodoService(*todoRepo)
// 	err = todoService.DeleteTodo(todo.ID)
// 	if err != nil {
// 		return
// 	}
// }
