package handlers

import (
	"encoding/json"
	"net/http"
	"roomserve/config"
	"roomserve/database"
	"roomserve/models"
	"roomserve/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func LoginUser(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	// parse json
	reqBody := new(models.LoginUser)
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusNotAcceptable)
		return
	}

	// find user by email
	user := models.User{}
	err = db.Raw("SELECT * FROM users WHERE email = ?", reqBody.Email).Scan(&user).Error
	if err == gorm.ErrRecordNotFound {
		http.Error(res, "Email not found", http.StatusBadRequest)
		return
	}

	// check password against database
	if !utils.CheckPasswordHash(reqBody.Password, user.Password) {
		http.Error(res, "Wrong credentials", http.StatusBadRequest)
		return
	}

	// create jwt token
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJson(res, 200, map[string]string{"token": token})
}
