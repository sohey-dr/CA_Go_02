package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"CA_Go/controller"
)

func NewRouter() error {
	r := mux.NewRouter()

	userSubRouter := r.PathPrefix("/user").Subrouter()
	userSubRouter.HandleFunc("/create", controller.CreateUser).Methods("POST")
	userSubRouter.HandleFunc("/get", controller.GetUser).Methods("GET")
	userSubRouter.HandleFunc("/update", controller.UpdateUser).Methods("PUT")

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), r)
}
