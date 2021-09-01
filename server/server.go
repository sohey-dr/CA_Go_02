package server

import (
	"CA_Go/database"
	"CA_Go/router"
)

func Run() {
	database.DbConnect()
	defer database.DB.Close()
	router.NewRouter()
}
