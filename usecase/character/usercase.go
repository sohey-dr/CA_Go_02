package character

import (
	"strconv"

	"CA_Go/model/character"
	"CA_Go/model/user"
	"CA_Go/model/user_character"
)

func GetUserCharacters(u *user.User) []user_character.UserCharacter {
	var characters []user_character.UserCharacter
	userCharacters := user_character.GetByUserID(strconv.FormatInt(u.ID, 10))
	for _, uc := range userCharacters {
		c := character.NewCharacter()
		c.FindById(uc.CharacterID)
		uc.Name = c.Name

		characters = append(characters, uc)
	}

	return characters
}
