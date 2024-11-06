package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Parse base template with the specified page content
	parsedTemplate, err := template.ParseFiles(
		"templates/index.html",
		filepath.Join("templates", tmpl),
	)

	// parse data

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parsedTemplate.Execute(w, data)
}

func RegistrationTemplateHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "registration.html", nil)
}

func UserListTemplateHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "list_data_user.html", nil)
}

// func TodoListTemplateHandler(w http.ResponseWriter, r *http.Request) {
// 	RenderTemplate(w, "todo_list.html")
// }

func UserDetailsTemplateHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "user_details.html", nil)
}
