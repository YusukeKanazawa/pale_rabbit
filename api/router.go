package api

import (
	"github.com/gorilla/mux"
	"net/http"
	//	"fmt"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

//var index = template.Must(template.ParseFiles(
//	"view/_base.html",
//	"view/index.html",
//))

var routes = Routes{
	Route{
		Method:      "GET",
		Pattern:     "/itemType",
		HandlerFunc: ItemTypeHandler,
	},
	Route{
		Method:      "PUT",
		Pattern:     "/itemType",
		HandlerFunc: ItemTypePutHandler,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
//		router.Methods(route.Method).Path(route.Pattern).Handler(route.HandlerFunc)
		router.Methods(route.Method).PathPrefix("/api").Path(route.Pattern).Handler(route.HandlerFunc)
	}
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	return router
}
