package handlers

import (
	"GOLANG-AUTH-SYSTEM/internal/db"
	"GOLANG-AUTH-SYSTEM/internal/models"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type RegisterReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {

	var req RegisterReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON Type", http.StatusBadRequest)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword( // Hash Generate
		[]byte(req.Password), bcrypt.DefaultCost,
	)

	if err != nil {
		http.Error(w, "Password Hash Failed", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPass),
	}

	err = db.DB.Create(&user).Error
	if err != nil {
		http.Error(w, "User registration failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})

	//fmt.Fprintln(w, req.Name)
}
