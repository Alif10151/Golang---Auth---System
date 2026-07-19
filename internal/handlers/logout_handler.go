package handlers

import (
	"encoding/json"
	"net/http"
)

func LogOut(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout Successful",
	})
}
