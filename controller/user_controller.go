package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"encoding/base64"
	"github.com/jinzhu/gorm"

	"CA_Go/model"
)

type TokenJson struct {
	Token string `json:"token"`
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

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
	
	token := auth.GenerateToken()

	user := model.NewUser()
	user.Name = data.Name
	user.Token = token
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	db.NewRecord(user)
	db.Create(&user)


	tokenJson := TokenJson{}
	tokenJson.Token = token
	outputJson, err := json.Marshal(&tokenJson)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(outputJson))
}