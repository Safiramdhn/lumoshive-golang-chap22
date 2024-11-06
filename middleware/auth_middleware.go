package middleware

import (
	"golang-beginner-22/database"
	"golang-beginner-22/repositories"
	"golang-beginner-22/services"
	"golang-beginner-22/utils"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := database.InitDB()
		if err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil || cookie == nil {
			// utils.RespondWithJSON(w, http.StatusUnauthorized, "Unauthorized", nil)
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
		}

		repo := repositories.NewUserRepositoryDB(db)
		service := services.NewUserService(*repo)
		token, err := service.GetUserByToken(cookie.Value)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		if token == "" || token != cookie.Value {
			// utils.RespondWithJSON(w, http.StatusUnauthorized, "Unauthorized", nil)
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
		}

		// Melanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
