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
	Get(string, HTTPFunc) HTTPRoute
	Post(string, HTTPFunc) HTTPRoute
	Put(string, HTTPFunc) HTTPRoute
	Patch(string, HTTPFunc) HTTPRoute
	Delete(string, HTTPFunc) HTTPRoute

	// returns all the routes
	GetRoutes() []HTTPRoute

	//register another router with the current router
	Handle(string, HTTPRouter)

	//register the router with a mux to handle http transport
	RegisterServer(*http.ServeMux)

	//called before the router is attached to the server
	Register()

	//registers a set of middlewares for the routers
	//applied to every route registered within the router
	//takes precedence over the middleware within the route
	Use(...MiddleWareFunc) HTTPRouter
}

// router contains a group of routes
// should implement the HTTPRouter interface
type Router struct {
	path   string
	routes []HTTPRoute

	middlewares []MiddleWareFunc
}

func NewRouter(path string) *Router {
	return &Router{
		path: strings.TrimRight(path, "/"),
	}
}

func (r *Router) createRoute(method HTTPMethod, path string, httpFunc HTTPFunc) HTTPRoute {
	path = strings.TrimRight(path, "/")
	route, err := NewRoute(method, r.path+path, httpFunc)
	if err != nil {
		slog.Error("Invalid route", "path", path)
		return nil
	}
	r.routes = append(r.routes, route)

	return route
}

func (r *Router) Get(path string, httpFunc HTTPFunc) HTTPRoute {
	return r.createRoute(GET, path, httpFunc)
}

func (r *Router) Put(path string, httpFunc HTTPFunc) HTTPRoute {
	return r.createRoute(PUT, path, httpFunc)
}

func (r *Router) Post(path string, httpFunc HTTPFunc) HTTPRoute {
	return r.createRoute(POST, path, httpFunc)
}

func (r *Router) Patch(path string, httpFunc HTTPFunc) HTTPRoute {
	return r.createRoute(PATCH, path, httpFunc)
}

func (r *Router) Delete(path string, httpFunc HTTPFunc) HTTPRoute {
	return r.createRoute(DELETE, path, httpFunc)
}

func (r *Router) GetRoutes() []HTTPRoute {
	return r.routes
}

// Registers a router with the given path
func (r *Router) Handle(path string, router HTTPRouter) {
	//calls the initialization of the router
	router.Register()

	r.path = strings.TrimRight(r.path, "/")
	r.path = fmt.Sprintf("%s%s", r.path, path)

	for _, route := range router.GetRoutes() {
		route.Prepend(r.path)
		route.stackMiddleware(r.middlewares)
		r.routes = append(r.routes, route)
	}
}

// empty function i dont know why but should be there
func (r *Router) Register() {
}

func (r *Router) RegisterServer(mux *http.ServeMux) {
	for _, route := range r.routes {
		route.stackMiddleware(r.middlewares)
		mux.HandleFunc(route.GetRoute(), route.GetHandler())
	}
}

// middleware stuff
// registers the middleware for entire router
func (r *Router) Use(middlewares ...MiddleWareFunc) HTTPRouter {
	r.middlewares = append(r.middlewares, middlewares...)
	return r
}
