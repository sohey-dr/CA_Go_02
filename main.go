package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/user/create", userCreate)
	
	http.ListenAndServe(":8080", nil)
}

func userCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go.")
}