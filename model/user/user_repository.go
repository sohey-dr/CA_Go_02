package user

import (
	"CA_Go/database"
)

func newUser() User {
	return User{}
}

func Create(params) {
	user := newUser()

}

func Find() {

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
