package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HTTPMethod string

// allowed methods
const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	PATCH  HTTPMethod = "PATCH"
	DELETE HTTPMethod = "DELETE"
)

// handles the URI of the api which is to be registered with the Router
type Route struct {
	method   HTTPMethod
	path     string
	httpfunc HTTPFunc
}

func (route *Route) GetRoute() string {
	return fmt.Sprintf("%s %s", route.method, route.path)
}

// return the http handler for the routes
// handles the encoding (json,grpc...)
func (route *Route) GetHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := route.httpfunc(w, r)

		//dont ask y coz i don't
		if err != nil {
			w.WriteHeader(err.code)
			json.NewEncoder(w).Encode(err)
			return
		} else if data != nil {
			w.WriteHeader(data.code)
			json.NewEncoder(w).Encode(data)
		}
	}
}

func NewRoute(method HTTPMethod, path string, httpfunc HTTPFunc) (*Route, error) {
	if len(path) == 0 || path[0] != '/' {
		return nil, errors.New("Invalid Route")
	}
	return &Route{
		method:   method,
		path:     path,
		httpfunc: httpfunc,
	}, nil
}
