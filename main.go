package main

import (
	"fmt"
	"golang-beginner-22/handlers"
	"golang-beginner-22/middleware"
	"html/template"
	"net/http"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Parse base template with the specified page content
	parsedTemplate, err := template.ParseFiles(
		"templates/base.html",
		filepath.Join("templates", tmpl),
	)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	parsedTemplate.Execute(w, nil)
}

func main() {
	serverMux := http.NewServeMux()

	// Register routes
	authMux := http.NewServeMux()
	authMux.HandleFunc("/register", handlers.RegistrationTemplateHandler)
	authMux.HandleFunc("/create-user", handlers.CreateUserHandler)

	userMux := http.NewServeMux()
	userMux.HandleFunc("/user-list", handlers.GetAllUsersHandler)
	userMux.HandleFunc("/user-detail", handlers.GetUserByIDHandler)
	userMiddleware := middleware.Middleware(userMux)
	serverMux.Handle("/user/", http.StripPrefix("/user", userMiddleware))

	todoMux := http.NewServeMux()
	todoMux.HandleFunc("/todo-list", handlers.GetTodosHandler)
	todoMux.HandleFunc("/create-todo", handlers.CreateTodoHandler)
	todoMux.HandleFunc("/update-todo", handlers.UpdateTodoHandler)
	todoMiddleware := middleware.Middleware(todoMux)

	serverMux.Handle("/", authMux)
	serverMux.Handle("/todo/", http.StripPrefix("/todo", todoMiddleware))

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", serverMux); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

	http.ListenAndServe(":8080", nil)
}
