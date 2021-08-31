package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Token string `json:"-"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type GetUserResponse struct {
	Name string `json:"name"`
}
