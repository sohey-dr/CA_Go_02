package user

import (
	"CA_Go/database"
	"math/rand"
	"time"
)

func NewUser() *User {
	return &User{}
}

func (u *User) Create(name string) {
	token := generateToken()

	u.Name = name
	u.Token = token
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	database.DB.NewRecord(u)
	database.DB.Create(&u)

	return
}

func (u *User) FindByToken(xToken string) {
	database.DB.Where("token = ?", xToken).First(&u)
}

func Update() {

}

func generateToken() string {
	codeAlphabet := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 20)
	for i := range b {
		b[i] = codeAlphabet[rand.Intn(len(codeAlphabet))]
	}
	return string(b)
}
