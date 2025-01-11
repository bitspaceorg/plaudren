package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type of the function that handles the http request
type HTTPFunc func(http.ResponseWriter, *http.Request) *ApiError

type HTTPMethod string

// type of the router which registers the functions
type HTTPRouter interface {

	// http methods
	// the param func should implement the HTTPFunc interface
	Get(string, HTTPFunc)
	Post(string, HTTPFunc)
	Put(string, HTTPFunc)
	Patch(string, HTTPFunc)
	Delete(string, HTTPFunc)

	// returns all the routes
	GetRoutes() []Route

	//register another router with the current router
	Handle(string, HTTPRouter)

	//register the router with a mux to handle http transport
	Register(*http.ServeMux)
}

// allowed methods
const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	PATCH  HTTPMethod = "PATCH"
	DELETE HTTPMethod = "DELETE"
)

/*
TODO:
1. fix strip '/' in url
2. test for nested routers
2. test for other methods
*/


//handles the URI of the api which is to be registered with the Router
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
		if err := route.httpfunc(w, r); err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
	}
}

func NewRoute(method HTTPMethod, path string, httpfunc HTTPFunc) *Route {
	return &Route{
		method:   method,
		path:     path,
		httpfunc: httpfunc,
	}
}

// router contains a group of routes
// should implement the HTTPRouter interface
type Router struct {
	path   string
	routes []Route
}

func NewRouter(path string) *Router {
	return &Router{
		path: path,
	}
}

func (r *Router) Get(path string, httpFunc HTTPFunc) {
	r.routes = append(r.routes, *NewRoute(GET, path, httpFunc))
}

func (r *Router) Put(path string, httpFunc HTTPFunc) {
	r.routes = append(r.routes, *NewRoute(PUT, path, httpFunc))
}

func (r *Router) Post(path string, httpFunc HTTPFunc) {
	r.routes = append(r.routes, *NewRoute(POST, path, httpFunc))
}

func (r *Router) Patch(path string, httpFunc HTTPFunc) {
	r.routes = append(r.routes, *NewRoute(PATCH, path, httpFunc))
}

func (r *Router) Delete(path string, httpFunc HTTPFunc) {
	r.routes = append(r.routes, *NewRoute(DELETE, path, httpFunc))
}

func (r *Router) GetRoutes() []Route {
	return r.routes
}

func (r *Router) Handle(path string, router HTTPRouter) {
	for _, route := range router.GetRoutes() {
		route.path = fmt.Sprintf("%s/%s", r.path, route.path)
		r.routes = append(r.routes, route)
	}
}

func (r *Router) Register(mux *http.ServeMux) {
	for _, route := range r.routes {
		mux.HandleFunc(route.GetRoute(), route.GetHandler())
	}
}
