package handlers

import (
	"GOLANG-AUTH-SYSTEM/internal/db"
	"GOLANG-AUTH-SYSTEM/internal/models"
	"encoding/json"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {

	email := r.Context().Value("email").(string)

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
