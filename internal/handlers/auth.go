package handlers

import (
	"GOLANG-AUTH-SYSTEM/internal/db"
	"GOLANG-AUTH-SYSTEM/internal/models"
	"GOLANG-AUTH-SYSTEM/internal/utils"
	"encoding/json"
	"net/http"
	"net/mail"
	"strings"

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
		utils.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	//fmt.Println(req.Email)

	var user models.User

	err = db.DB.Where("email=?", req.Email).First(&user).Error
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, "Invalid Email")
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		utils.Error(w, http.StatusUnauthorized, "Invalid Password")
		return
	}

	// json.NewEncoder(w).Encode(map[string]string{
	// 	"message": "Login Successful",
	// })

	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	w.Header().Set("Content-Type", "Application/json")

	utils.Success(w, http.StatusOK, "Login Successful", map[string]string{
		"token": token,
	},
	)

}

func Register(w http.ResponseWriter, r *http.Request) {

	var req RegisterReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid JSON Type")
		return
	}

	// Remove extra spaces
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	//validation
	if strings.TrimSpace(req.Name) == "" {
		utils.Error(w, http.StatusBadRequest, "Name is required")
		return
	}

	if strings.TrimSpace(req.Email) == "" {
		utils.Error(w, http.StatusBadRequest, "Email is required")
		return
	}

	if strings.TrimSpace(req.Password) == "" {
		utils.Error(w, http.StatusBadRequest, "Password is required")
		return
	}

	if len(req.Password) < 6 {
		utils.Error(w, http.StatusBadRequest, "Password must be at least 6 digits")
		return
	}

	// Email format validation
	_, err = mail.ParseAddress(req.Email)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid Email Formet")
		return
	}

	var existingUser models.User

	err = db.DB.Where("email=?", req.Email).First(&existingUser).Error
	if err == nil { // checking the email
		utils.Error(w, http.StatusConflict, "Email already existed")
		return // email exists already , so return the process of register
	}

	// else do register and hash pass
	hashedPass, err := bcrypt.GenerateFromPassword( // Hash Generate
		[]byte(req.Password), bcrypt.DefaultCost,
	)

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Password Hash Failed")
		return
	}

	//user create
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPass),
	}

	err = db.DB.Create(&user).Error
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "User Registration Failed")
		return
	}

	utils.Success(w, http.StatusCreated, "User Registered Successfully", nil)
}
