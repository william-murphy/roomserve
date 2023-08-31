package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"roomserve/config"
	"roomserve/database"
	"roomserve/models"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func Protected(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		db := database.DB

		// get token string from headers
		tokenStr := req.Header.Get("Authorization")
		splitToken := strings.Split(tokenStr, "Bearer ")
		if len(splitToken) != 2 {
			http.Error(res, "Missing or malformed JWT", http.StatusUnauthorized)
			return
		}
		tokenStr = splitToken[1]

		// decode token string
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("incorrect signing method")
			}
			return []byte(config.Config("SECRET")), nil
		})
		if err != nil || !token.Valid { // TODO - do i manually check if token is expired?
			fmt.Println(err.Error())
			http.Error(res, "Invalid or expired JWT", http.StatusUnauthorized)
			return
		}

		// get user from id in token
		var user models.User
		id := token.Claims.(jwt.MapClaims)["id"]
		err = db.First(&user, id).Error
		if err == gorm.ErrRecordNotFound {
			http.Error(res, "JWT represents nonexistent user", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), "user", user)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
