package api

import (
	"fmt"
	"net/http"
//	"regexp"
)

//var route_list []*Route = []*Route{
//	&Route{
//		Method: "GET",
//		Path:   "/api/items/{code}",
//	},
//}

func Router(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI
//	for _, route := range route_list{
//		route.Path
//	}

	fmt.Fprint(w, path)
}

//type Route struct {
//	Method  string
//	Path    string
//	Handler http.Handler
//}
