package utils

import (
	"encoding/json"
	"net/http"
	"regexp"
	"roomserve/database"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func RespondWithJson(res http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)
	res.Write(response)
}

// CheckPasswordHash compares password with hash
func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

func CheckOverlappingTime(id uint, start time.Time, end time.Time, roomId uint) bool {
	db := database.DB
	var Found bool
	db.Raw("SELECT EXISTS(SELECT 1 FROM reservations WHERE reservations.id != ? AND reservations.start <= ? AND reservations.end >= ? AND reservations.room_id = ?) AS found", id, end, start, roomId).Scan(&Found)
	return Found
}

func ExceedsRoomCapacity(numUsers int, roomId uint) bool {
	db := database.DB
	var Found bool
	db.Raw("SELECT true AS found FROM rooms WHERE id = ? AND capacity < ?", roomId, numUsers).Scan(&Found)
	return Found
}

func ConvertSearchQuery(query string) string {
	re := regexp.MustCompile("[^A-Za-z0-9]+")
	splitQuery := re.Split(query, -1)
	result := strings.Join(splitQuery, " | ")
	return result
}
