package gacha

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"CA_Go/usercase/gacha"
)

func DrawCharacter(w http.ResponseWriter, r *http.Request) {
	params := &gacha.GachaDrawRequest{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		fmt.Println("JSON Decode error:", err)
		return
	}

	xToken := r.Header.Get("x-token")

	results := gacha.Draws(params.Times, xToken)

	response := gacha.GachaDrawResponse{Results: results}
	j, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(j)))
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
