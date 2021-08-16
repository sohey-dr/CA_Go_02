package main

import (
	"CA_Go/database"
	"CA_Go/router"
)

func main() {
	// ここ要らない
	// routerを作る
	database.DbConnect()
	defer database.DB.Close()
	router.NewRouter()
}
