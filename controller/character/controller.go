package character

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CA_Go/model/user"
	"CA_Go/model/user_character"
	"CA_Go/usercase/character"
)

func GetUserCharacters(w http.ResponseWriter, r *http.Request) {
	xToken := r.Header.Get("x-token")

	u := user.NewUser()
	u.FindByToken(xToken)

	characters := character.GetUserCharacters(u)

	response := &user_character.CharacterListResponse{Characters: characters}
	j, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(j)))
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
