package handlers

import (
	"encoding/json"
	"fmt"
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

	fmt.Println(string(hashedPass))
	fmt.Fprintln(w, req.Name)
}
