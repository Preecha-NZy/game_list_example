package main

import "net/http"

func main() {
	router := makeRoute()
	http.Handle("/", router)
	http.ListenAndServe(":8000", router)
}
