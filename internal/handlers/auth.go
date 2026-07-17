package handlers

import (
	"GOLANG-AUTH-SYSTEM/internal/db"
	"GOLANG-AUTH-SYSTEM/internal/models"
	"GOLANG-AUTH-SYSTEM/internal/utils"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type RegisterReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//fmt.Println(req.Email)

	var user models.User

	err = db.DB.Where("email=?", req.Email).First(&user).Error
	if err != nil {
		http.Error(w, "Invalid Email", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		return
	}

	// json.NewEncoder(w).Encode(map[string]string{
	// 	"message": "Login Successful",
	// })

	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "Application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login Successful",
		"token":   token,
	})

}

func Register(w http.ResponseWriter, r *http.Request) {

	var req RegisterReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON Type", http.StatusBadRequest)
		return
	}

	var existingUser models.User

	err = db.DB.Where("email=?", req.Email).First(&existingUser).Error
	if err == nil { // checking the email
		http.Error(w, "Email already existed", http.StatusConflict)
		return // email exists already , so return the process of register
	}

	// else do register and hash pass
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
