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
	"golang.org/x/crypto/bcrypt"
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
	db.Raw("SELECT * FROM users WHERE email = ?", reqBody.Email).Scan(&user)
	if user.ID < 1 {
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

func RegisterUser(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	// parse json
	reqBody := new(models.RegisterUser)
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusNotAcceptable)
		return
	}

	// create a hash of the given password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(res, "Unable to register user", http.StatusNotAcceptable)
		return
	}

	// create user (with hashed password)
	newUser := models.User{
		Name:     reqBody.Name,
		Username: reqBody.Username,
		Email:    reqBody.Email,
		Password: hashedPassword,
	}
	err = db.Create(&newUser).Error
	if err != nil {
		http.Error(res, "Unable to register user", http.StatusNotAcceptable)
		return
	}

	utils.RespondWithJson(res, 201, newUser)
}
