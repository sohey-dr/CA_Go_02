package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/jinzhu/gorm"

	"CA_Go/model"
	"CA_Go/auth"
)

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

	response := model.CreateUserResponse{}
	response.Token = token
	outputJson, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(outputJson))
}

func GetUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != "GET" {
		fmt.Fprint(w, "Not get...")
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
	fmt.Fprint(w, string(outputJson))
}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set( "Access-Control-Allow-Methods","PUT" )

	if r.Method != "PUT" {
		fmt.Fprint(w, "Not put...")
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

	fmt.Fprint(w, "A successful response.")
}