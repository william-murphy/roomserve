package middleware

import (
	"net/http"
)

func Json(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		//check if request is application/json
		contentType := req.Header.Get("Content-Type")
		if contentType != "" && contentType != "application/json" {
			http.Error(res, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
			return
		}

		// add accepts application/json header to response
		res.Header().Set("Accepts", "application/json")

		next.ServeHTTP(res, req)
	})
}
