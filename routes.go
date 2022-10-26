package main

import (
	"gle/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func makeRoute() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/games", controller.Games{}.Index).Methods("Get")
	return r
}
