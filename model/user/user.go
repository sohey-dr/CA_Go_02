package user

import (
	"CA_Go/model/character"
	"time"
)

type User struct {
	ID         int64
	Name       string                `json:"name"`
	Token      string                `json:"-"`
	Characters []character.Character `gorm:"many2many:user_characters"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type GetUserResponse struct {
	Name string `json:"name"`
}
