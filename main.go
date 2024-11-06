package main

import (
	"fmt"
	"golang-beginner-22/handlers"
	"golang-beginner-22/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", handlers.RegisterHandler)
			r.Post("/login", handlers.LoginHandler)

		})

		r.Route("/todos", func(r chi.Router) {
			r.With(middleware.AuthMiddleware).Post("/create", handlers.CreateTodoHandler)
			r.With(middleware.AuthMiddleware).Post("/update/{id}", handlers.UpdateTodoHandler)
		})
	})

	r.Get("/register", handlers.RegisterFormHandler)
	r.Get("/login", handlers.LoginFormHandler)
	r.With(middleware.AuthMiddleware).Get("/user-list", handlers.UserListHandler)
	r.With(middleware.AuthMiddleware).Get("/todo-list", handlers.UserTodoListHandler)
	r.With(middleware.AuthMiddleware).Get("/user-detail/{id}", handlers.UserDetailHandler)

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

	http.ListenAndServe(":8080", r)
}
