package main

import (
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
	http.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "registration.html")
	})
	http.HandleFunc("/list-data-user", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "list_data_user.html")
	})
	http.HandleFunc("/todo-list", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "todo_list.html")
	})
	http.HandleFunc("/user-details", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "user_details.html")
	})

	http.ListenAndServe(":8080", nil)
}
