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

func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
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
	
	token := base64.StdEncoding.EncodeToString([]byte(data.Name))

	user := model.NewUser()
	user.Name = data.Name
	user.Token = token
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()


	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(body))
}