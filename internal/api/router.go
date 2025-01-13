package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

// type of the function that handles the http request
type HTTPFunc func(http.ResponseWriter, *http.Request) (*ApiData, *ApiError)

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
	RegisterServer(*http.ServeMux)

	//called before the router is attached to the server
	Register()
}

// router contains a group of routes
// should implement the HTTPRouter interface
type Router struct {
	path   string
	routes []Route
}

func NewRouter(path string) *Router {
	return &Router{
		path: strings.TrimRight(path, "/"),
	}
}

func (r *Router) createRoute(method HTTPMethod, path string, httpFunc HTTPFunc) {
	path = strings.TrimRight(path, "/")
	route, err := NewRoute(method, r.path+path, httpFunc)
	if err != nil {
		slog.Error("Invalid route", "path", path)
		return
	}
	r.routes = append(r.routes, *route)
}

func (r *Router) Get(path string, httpFunc HTTPFunc) {
	r.createRoute(GET, path, httpFunc)
}

func (r *Router) Put(path string, httpFunc HTTPFunc) {
	r.createRoute(PUT, path, httpFunc)
}

func (r *Router) Post(path string, httpFunc HTTPFunc) {
	r.createRoute(POST, path, httpFunc)
}

func (r *Router) Patch(path string, httpFunc HTTPFunc) {
	r.createRoute(PATCH, path, httpFunc)
}

func (r *Router) Delete(path string, httpFunc HTTPFunc) {
	r.createRoute(DELETE, path, httpFunc)
}

func (r *Router) GetRoutes() []Route {
	return r.routes
}

// Registers a router with the given path
func (r *Router) Handle(path string, router HTTPRouter) {
	//calls the initialization of the router
	router.Register()

	r.path = strings.TrimRight(r.path, "/")
	r.path = fmt.Sprintf("%s%s", r.path, path)

	for _, route := range router.GetRoutes() {
		route.path = fmt.Sprintf("%s%s", r.path, strings.TrimRight(route.path, "/"))
		r.routes = append(r.routes, route)
	}
}

// empty function i dont know why but should be there
func (r *Router) Register() {
}

func (r *Router) RegisterServer(mux *http.ServeMux) {
	for _, route := range r.routes {
		mux.HandleFunc(route.GetRoute(), route.GetHandler())
	}
}
