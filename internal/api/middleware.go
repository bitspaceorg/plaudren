package api

import "net/http"

// the type of middleware function
// the routers or the routes themself can have middlewares
// the middleware registered in the router takes precendence over the middleware registered in the routes
// if a error is returned the middleware chain is terminated , else the next middleware or the function is automatically called
type MiddleWareFunc func(http.ResponseWriter, *http.Request) *ApiError
