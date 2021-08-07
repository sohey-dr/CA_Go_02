package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"CA_Go/model"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
			fmt.Fprint(w, "Not post...")
			return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
			fmt.Println("io error")
			return
	}

	jsonBytes := ([]byte)(body)
	data := new(model.User)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
			fmt.Println("JSON Unmarshal error:", err)
			return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(body))
}