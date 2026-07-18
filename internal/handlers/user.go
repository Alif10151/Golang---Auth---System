package handlers

import (
	"GOLANG-AUTH-SYSTEM/internal/db"
	"GOLANG-AUTH-SYSTEM/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ProfileResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Profile(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(map[string]string{
	// 	"message": "Welcome to Profile",
	// })
	email := r.Context().Value("email").(string)
	//fmt.Println(email)

	var user models.User

	err := db.DB.Where("email=?", email).First(&user).Error

	if err != nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	response := ProfileResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	json.NewEncoder(w).Encode(response)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	err := db.DB.Find(&users).Error

	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	var response []models.UserResponse

	for _, user := range users {
		response = append(response, models.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	json.NewEncoder(w).Encode(response)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/users/")
	userID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	var user models.User
	err = db.DB.First(&user, userID).Error

	if err != nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	response := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	json.NewEncoder(w).Encode(response)

}
