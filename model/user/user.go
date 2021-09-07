package user

import (
	"github.com/jinzhu/gorm"
	"CA_Go/model/character"
)

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Token       string `json:"-"`
	Characters  []character.Character `gorm:"many2many:user_characters"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type GetUserResponse struct {
	Name string `json:"name"`
}
