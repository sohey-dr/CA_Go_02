package main

import (
	"net/http"

	"CA_Go/database"
	"CA_Go/model"
	"CA_Go/controller"
)

func main() {
	db := database.DbConnect()
	http.HandleFunc("/user/create", controller.UserCreate)
	defer db.Close()

	db.AutoMigrate(&model.User{})
	http.ListenAndServe(":8080", nil)
}