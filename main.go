package main

import (
	"net/http"
	"fmt"

	"CA_Go/database"
	"CA_Go/model"
)

func main() {
	db := database.DbConnect()
	http.HandleFunc("/user/create", userCreate)
	defer db.Close()

	db.AutoMigrate(&model.User{})
	http.ListenAndServe(":8080", nil)
}

func userCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go.")
}