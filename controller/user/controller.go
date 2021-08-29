package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"CA_Go/database"
	"CA_Go/model/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("io error")
		return
	}

	jsonBytes := ([]byte)(body)
	params := new(user.User)
	if err := json.Unmarshal(jsonBytes, params); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	u := user.NewUser()
	token := u.Create(params.Name)

	response := user.CreateUserResponse{}
	response.Token = token
	outputJson, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = fmt.Fprint(w, string(outputJson))
	if err != nil {
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != "GET" {
		_, err := fmt.Fprint(w, "Not get...")
		if err != nil {
			return
		}
		return
	}

	xToken := r.Header.Get("x-token")
	u := user.NewUser()

	database.DB.Where("token = ?", xToken).First(&u)

	response := user.GetUserResponse{}
	response.Name = u.Name
	outputJson, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = fmt.Fprintf(w, string(outputJson))
	if err != nil {
		return
	}
	return
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	if r.Method != "PUT" {
		_, err := fmt.Fprint(w, "Not put...")
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
	data := new(user.User)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	xToken := r.Header.Get("x-token")
	user := user.NewUser()

	database.DB.Where("token = ?", xToken).First(&user).Update("name", data.Name)
	if user.ID == 0 {
		fmt.Println("don't find user id:", user.ID)
		return
	}
	fmt.Println("updated user id:", user.ID)

	_, err = fmt.Fprint(w, "A successful response.")
	if err != nil {
		return
	}
}
