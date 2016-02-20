package main

import (
	"net/http"
	"pale_rabbit/api"
)

func main() {
	http.HandleFunc("/", api.Router)
	http.ListenAndServe("localhost:8080", nil)
}
