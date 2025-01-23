package api

import (
	"fmt"
	"net/http"
	"strings"
)

type FileHandler struct {
	handler     http.Handler
	path        string
	middlewares []MiddleWareFunc
}

func NewFileHandler(dir http.FileSystem, path string) *FileHandler {
	if path == "" {
		path = "/"
	}
	return &FileHandler{
		handler: http.FileServer(dir),
		path:    path,
	}
}

func (h *FileHandler) GetRoute() string {
	if h.path != "/" {
		return fmt.Sprintf("%s/", h.path)
	}
	return h.path
}

func (h *FileHandler) applyMiddleware(w http.ResponseWriter, r *http.Request) *ApiError {
	for _, middleware := range h.middlewares {
		if err := middleware(w, r); err != nil {
			return err
		}
	}
	return nil
}

func (h *FileHandler) GetHandleFunc() func(http.ResponseWriter, *http.Request) {
	return nil
}

func (h *FileHandler) GetHandler() http.Handler {
	return h.handler
}

func (h *FileHandler) stackMiddleware(middleware []MiddleWareFunc) {
	h.middlewares = append(middleware, h.middlewares...)
}

// registers a set of all middlewares
// adds the middlewares in order
func (h *FileHandler) Use(middlewares ...MiddleWareFunc) {
	h.middlewares = append(h.middlewares, middlewares...)
}

func (h *FileHandler) Prepend(path string) {
	h.path = fmt.Sprintf("%s%s", path, strings.TrimRight(h.path, "/"))
}
