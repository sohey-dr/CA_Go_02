package main

import (
	"net/http"

	"CA_Go/database"
	"CA_Go/controller"
)

func main() {
	db := database.DbConnect()

	// NOTE:HandleFuncを1行で書いて、第2引数のメソッドに第3引数を渡したい
	http.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
    controller.CreateUser(w, r, db)
  })
	http.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
    controller.GetUser(w, r, db)
  })

	defer db.Close()

	http.ListenAndServe(":8080", nil)
}