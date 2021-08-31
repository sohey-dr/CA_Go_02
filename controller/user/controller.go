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
		fmt.Println("JSON Decode error:", err)
		return
	}

	u := user.NewUser()
	u.Create(params.Name)

	w.Header().Set("Content-Type", "application/json")

	var b bytes.Buffer
	response := &user.CreateUserResponse{Token: u.Token}
	err = json.NewEncoder(&b).Encode(response)
	if err != nil {
		fmt.Println("JSON Encode error:", err)
		return
	}

	_, err = fmt.Fprint(w, b.String())
	if err != nil {
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	xToken := r.Header.Get("x-token")

	u := user.NewUser()
	u.FindByToken(xToken)

	w.Header().Set("Content-Type", "application/json")
	var b bytes.Buffer
	response := user.GetUserResponse{Name: u.Name}
	err = json.NewEncoder(&b).Encode(response)
	if err != nil {
		fmt.Println("JSON Encode error:", err)
		return
	}

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
