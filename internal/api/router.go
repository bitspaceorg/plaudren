package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Router struct {
	path   string
	router *http.ServeMux
}

type HTTPFunc func(http.ResponseWriter, *http.Request) *ApiError

func NewRouter(path string) *Router {
	return &Router{
		path:   path,
		router: &http.ServeMux{},
	}
}

func (r *Router) GetPath() string {
	return r.path
}

func (r *Router) GetHandler() http.Handler {
	return r.router
}

func (r *Router) Get(path string, httpFunc HTTPFunc) {
	r.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("type", "application/json")
		if err := httpFunc(w, r); err != nil {
			w.WriteHeader(err.code)
			err := json.NewEncoder(w).Encode(err)
			if err != nil {
				slog.Error(err.Error())
			}
		}
	})
}
