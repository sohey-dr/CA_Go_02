package controller

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"time"

	"CA_Go/auth"
	"CA_Go/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// handlerを分ける(ヘッダーセットやbody読み取る部分、レスポンス部分)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != "POST" {
		fprint, err := fmt.Fprint(w, "Not post...")
		if err != nil {
			return
		}
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

	response := model.CreateUserResponse{}
	response.Token = token
	outputJson, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fprint, err := fmt.Fprint(w, string(outputJson))
	if err != nil {
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != "GET" {
		fprint, err := fmt.Fprint(w, "Not get...")
		if err != nil {
			return
		}
		return
	}

	xToken := r.Header.Get("x-token")
	user := model.NewUser()

	db.Where("token = ?", xToken).First(&user)

	response := model.GetUserResponse{}
	response.Name = user.Name
	outputJson, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fprintf, err := fmt.Fprintf(w, string(outputJson))
	if err != nil {
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	if r.Method != "PUT" {
		fprint, err := fmt.Fprint(w, "Not put...")
		if err != nil {
			return
		}
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

	xToken := r.Header.Get("x-token")
	user := model.NewUser()

	db.Where("token = ?", xToken).First(&user).Update("name", data.Name)
	if user.ID == 0 {
		fmt.Println("don't find user id:", user.ID)
		return
	}
	fmt.Println("updated user id:", user.ID)

	fprint, err := fmt.Fprint(w, "A successful response.")
	if err != nil {
		return
	}
}
