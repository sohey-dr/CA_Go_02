package database

import (
	"CA_Go/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("db connected: ", &db)
	db.AutoMigrate(&model.User{})
	return db
}
