package middleware

import (
	"net/http"
	"roomserve/config"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		clientUrl := config.Config("CLIENT_URL")

		res.Header().Set("Access-Control-Allow-Origin", clientUrl)
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		res.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")

		if req.Method == http.MethodOptions {
			res.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(res, req)
	})
}
