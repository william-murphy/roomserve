package middleware

import (
	"net/http"
	"roomserve/models"
)

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		user := req.Context().Value("user").(models.User)
		if !user.IsAdmin {
			http.Error(res, "Admin access only", http.StatusForbidden)
			return
		}
		next.ServeHTTP(res, req)
	})
}
