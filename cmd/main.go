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

	mux := http.NewServeMux()

	http.HandleFunc("/checking_activity", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server Running Successfully")
	})

	mux.HandleFunc("/register", handlers.Register)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/profile", middleware.JWTAuth(handlers.Profile))
	mux.HandleFunc("/logout", handlers.LogOut)
	mux.HandleFunc("/users", handlers.GetUsers)
	mux.HandleFunc("/users/", handlers.GetUserByID)
	mux.HandleFunc("/update_user", handlers.UpdateUser)
	mux.HandleFunc("/delete_user", handlers.DeleteUser)

	handler := middleware.Logger(middleware.CORS(mux.ServeHTTP))

	fmt.Println("Server is running")
	http.ListenAndServe(":8800", handler)

}
