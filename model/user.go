package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Token string
}

func NewUser() User {
	return User{}
}