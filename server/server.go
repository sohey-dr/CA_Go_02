package server

import (
	"CA_Go/database"
	"CA_Go/router"
)

func Run() {
	database.DBConnect()
	defer database.DB.Close()
	router.NewRouter()
}
