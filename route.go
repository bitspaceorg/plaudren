package plaud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
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

// leaf node of a router
type HTTPRoute interface {
	GetRoute() string
	GetHandleFunc() func(http.ResponseWriter, *http.Request)
	GetHandler() http.Handler

	// registers a pre route from a router or handler
	Prepend(string)

	// registers the routers middlewares to the route
	stackMiddleware([]MiddleWareFunc)
	// registers route specific middleware
	Use(...MiddleWareFunc)
}

// handles the URI of the api which is to be registered with the Router
type Route struct {
	method      HTTPMethod
	path        string
	httpfunc    HTTPFunc
	middlewares []MiddleWareFunc
}

func (route *Route) GetRoute() string {
	return fmt.Sprintf("%s %s", route.method, route.path)
}

// applies a set of all middlware to a route
func (route *Route) applyMiddlware(w http.ResponseWriter, r *http.Request) *Error {
	for _, middleware := range route.middlewares {
		if err := middleware(w, r); err != nil {
			return err
		}
	}
	return nil
}

// return the http handler for the routes
// handles the encoding (json,grpc...)
//
//nolint:errcheck // TODO: Error handling will be added in a future commit
func (route *Route) GetHandleFunc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// handling middlewares
		if err := route.applyMiddlware(w, r); err != nil {
			w.WriteHeader(err.code)
			// TODO: handle error below
			// have a default logger with the router
			json.NewEncoder(w).Encode(err)

			return
		}
		data, err := route.httpfunc(w, r)
		// dont ask y coz i don't
		if err != nil {
			w.WriteHeader(err.code)
			// TODO: handle error below
			// have a default logger with the router
			json.NewEncoder(w).Encode(err)
		} else if data != nil {
			w.WriteHeader(data.code)
			// TODO: handle error below
			// have a default logger with the router
			json.NewEncoder(w).Encode(data)
		}
	}
}

func (route *Route) GetHandler() http.Handler {
	return nil
}
func NewRoute(method HTTPMethod, path string, httpfunc HTTPFunc) (*Route, error) {
	if path == "" {
		path = "/"
	}
	if path[0] != '/' {
		return nil, errors.New("invalid route")
	}
	return &Route{
		method:   method,
		path:     path,
		httpfunc: httpfunc,
	}, nil
}

func (route *Route) stackMiddleware(middleware []MiddleWareFunc) {
	route.middlewares = append(middleware, route.middlewares...)
}

// registers a set of all middlewares
// adds the middlewares in order
func (route *Route) Use(middlewares ...MiddleWareFunc) {
	route.middlewares = append(route.middlewares, middlewares...)
}

func (route *Route) Prepend(path string) {
	route.path = fmt.Sprintf("%s%s", path, strings.TrimRight(route.path, "/"))
}
