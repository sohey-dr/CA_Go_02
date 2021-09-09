package gacha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"CA_Go/model/user"
	//"CA_Go/model/character"
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
	results := gacha.Draws(2)
	fmt.Println(results[0])

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(j)))
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
