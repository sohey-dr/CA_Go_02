package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"CA_Go/model/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	params := &user.User{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		fmt.Println("JSON Decode error:", err)
		return
	}

	u := user.NewUser()
	u.Create(params.Name)

	response := &user.CreateUserResponse{Token: u.Token}
	j, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(j)))
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	xToken := r.Header.Get("x-token")

	u := user.NewUser()
	u.FindByToken(xToken)

	response := user.GetUserResponse{Name: u.Name}
	j, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(j)))
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := &user.User{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		fmt.Println("JSON Decode error:", err)
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
