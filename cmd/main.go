package main

import (
	"GOLANG-AUTH-SYSTEM/internal/db"
	"GOLANG-AUTH-SYSTEM/internal/handlers"
	"fmt"
	"net/http"
)

func main() {

	db.ConnectDB()

	http.HandleFunc("/checking_activity", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server Running Successfully")
	})

	http.HandleFunc("/register", handlers.Register)

	fmt.Println("Server is running")
	http.ListenAndServe(":8800", nil)

}
