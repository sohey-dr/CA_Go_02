package user_character

import (
	"CA_Go/database"
)

func GetByUserID(userId string) []UserCharacter {
	var userCharacters []UserCharacter
	database.DB.Where("user_id = ?", userId).Find(&userCharacters)
	return userCharacters
}
