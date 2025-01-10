package api

import (
	"log/slog"
	"net/http"
)

type ApiHandler interface {
	GetPath() string
	GetHandler() http.Handler
	Register()
}

type ApiServer struct {
	listenAddr string
	server     *http.ServeMux
}

func New(listenAddr string) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		server:     http.NewServeMux(),
	}
}

func (s *ApiServer) Register(routers ...ApiHandler) {
	for _, router := range routers {
		router.Register()
		s.server.Handle(router.GetPath()+"/", http.StripPrefix(router.GetPath(), router.GetHandler()))
	}
}

func (s *ApiServer) Run() error {
	slog.Info("Server started on", "port", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, s.server)
}
