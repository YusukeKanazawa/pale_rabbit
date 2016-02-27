package main

import (
	"log"
	"net/http"
	"os"
	"pale_rabbit/api"
)

const (
	DEFAULT_PORT = "9000"
	DEFAULT_HOST = "localhost"
)

func main() {
	var port string
	if port = os.Getenv("VCAP_APP_PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}
	var host string
	if host = os.Getenv("VCAP_APP_HOST"); len(host) == 0 {
		host = DEFAULT_HOST
	}
	//	app.Init()
	router := api.NewRouter()
	http.Handle("/", router)
	log.Printf("Starting app on %+v:%+v\n", host, port)
	http.ListenAndServe(host+":"+port, nil)
}
