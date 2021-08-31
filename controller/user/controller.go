package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"CA_Go/model/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	params := &user.User{}
	err := json.NewDecoder(r.Body).Decode(params)
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
	u.Create(params.Name)

	response := user.CreateUserResponse{Token: u.Token}
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
	xToken := r.Header.Get("x-token")

	u := user.NewUser()
	u.FindByToken(xToken)

	response := user.GetUserResponse{Name: u.Name}
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
	params := &user.User{}
	err := json.NewDecoder(r.Body).Decode(params)
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

	u := user.NewUser()
	u.UpdateNameByToken(xToken, params.Name)
	if u.ID == 0 {
		fmt.Println("don't find user id:", u.ID)
		return
	}
	fmt.Println("updated user id:", u.ID)

	_, err = fmt.Fprint(w, "A successful response.")
	if err != nil {
		return
	}
}
