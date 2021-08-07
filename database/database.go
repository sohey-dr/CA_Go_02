package database

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("PASSWORD")
	protcol := "tcp(" + os.Getenv("DOMAIN") + ":" + os.Getenv("PORT") + ")"
	dbName := os.Getenv("DB_NAME") + "?charset=utf8&parseTime=true&loc=Local"
	CONNECT := dbUser + ":" + dbPass + "@" + protcol + "/" + dbName

	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("db connected: ", &db)
	return db
}