package middleware

import (
	"golang-beginner-22/database"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("token")
		queryStatement := `SELECT id FROM users WHERE token = $1`
		var userID int
		err := db.QueryRow(queryStatement, authHeader).Scan(userID)
		if err != nil || userID == 0 {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			// utils.RespondWithJSON(w, http.StatusUnauthorized, "Unauthorized", nil)
			return
		}

		// Melanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
