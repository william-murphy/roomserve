package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

func GetUserIdFromToken(token *jwt.Token) uint {
	return uint(token.Claims.(jwt.MapClaims)["id"].(float64))
}
