package main

import (
	//  mux
	"net/http"

	"CA_Go/controller"
	"CA_Go/database"
)

func main() {
	// ここ要らない
	// routerを作る
	database.DbConnect()
	defer database.DB.Close()

	// NOTE:HandleFuncを1行で書いて、第2引数のメソッドに第3引数を渡したい
	http.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		controller.CreateUser(w, r)
	})
	http.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
		controller.GetUser(w, r)
	})
	http.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		controller.UpdateUser(w, r)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
