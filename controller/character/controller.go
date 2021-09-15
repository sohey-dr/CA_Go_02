package character

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CA_Go/model/character"
	"CA_Go/model/user"
	"CA_Go/model/user_character"
)

func GetUserCharacters(w http.ResponseWriter, r *http.Request) {
	xToken := r.Header.Get("x-token")

	u := user.NewUser()
	u.FindByToken(xToken)

	var characters []user_character.UserCharacter
	userCharacters := user_character.GetByUserID(strconv.FormatInt(u.ID, 10))
	for _, uc := range userCharacters {
		c := character.NewCharacter()
		c.FindById(uc.CharacterID)
		uc.Name = c.Name

		characters = append(characters, uc)
	}

	response := &user_character.CharacterListResponse{Characters: characters}
	j, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(j)))
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
