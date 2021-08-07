package database

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
)

func DbConnect() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("DOMAIN") + ":" + os.Getenv("PORT") + ")"
	DBNAME := os.Getenv("DBNAME") + "?parseTime=true&loc=Asia%2FTokyo"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("db connected: ", &db)
	return db
}