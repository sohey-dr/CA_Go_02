package main

import (
	"CA_Go/database"
	"CA_Go/router"
)

func main() {
	database.DbConnect()
	defer database.DB.Close()
	router.NewRouter()
}
