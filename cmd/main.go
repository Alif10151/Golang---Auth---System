package main

import (
	"GOLANG-AUTH-SYSTEM/internal/db"
	"GOLANG-AUTH-SYSTEM/internal/handlers"
	"GOLANG-AUTH-SYSTEM/internal/middleware"
	"fmt"
	"net/http"
)

func main() {

	db.ConnectDB()

	http.HandleFunc("/checking_activity", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server Running Successfully")
	})

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/profile", middleware.JWTAuth(handlers.Profile))
	http.HandleFunc("/logout", handlers.LogOut)
	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/users/", handlers.GetUserByID)
	http.HandleFunc("/update_user", handlers.UpdateUser)

	fmt.Println("Server is running")
	http.ListenAndServe(":8800", nil)

}
