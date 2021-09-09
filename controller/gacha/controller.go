package gacha

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CA_Go/model/user"
	"CA_Go/usercase/gacha"
)

func DrawCharacter(w http.ResponseWriter, r *http.Request) {
	xToken := r.Header.Get("x-token")

	u := user.NewUser()
	u.FindByToken(xToken)

	//characters := []character.Character{{}, {}}
	//db.Create(&characters)
	//user := user.User{Character: characters}
	//db.Create(&user)
	characters := gacha.Draws(2)

	var results []gacha.GachaResult
	for _, character := range characters {
		result := gacha.GachaResult{CharacterID: strconv.Itoa(character.ID), Name: character.Name}
		results = append(results, result)
	}

	response := gacha.GachaDrawResponse{Results: results}
	j, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(j)))
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
