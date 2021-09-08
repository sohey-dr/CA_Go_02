package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

var DB *gorm.DB

func DbConnect() {
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
	DB = db

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("db connected: ", &DB)
	//TODO:どこかでマイグレーションする
}
