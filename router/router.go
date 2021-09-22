package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"

	"CA_Go/controller/character"
	"CA_Go/controller/gacha"
	"CA_Go/controller/user"
)

func NewRouter() error {
	r := mux.NewRouter()
	userSubRouter := r.PathPrefix("/user").Subrouter()

	userSubRouter.HandleFunc("/create", user.CreateUser).Methods("POST")
	userSubRouter.HandleFunc("/get", user.GetUser).Methods("GET")
	userSubRouter.HandleFunc("/update", user.UpdateUser).Methods("PUT")

	r.HandleFunc("/gacha/draw", gacha.DrawCharacter).Methods("POST")
	r.HandleFunc("/character/list", character.GetUserCharacters).Methods("GET")

	c := cors.AllowAll().Handler(r)
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), c)
}
