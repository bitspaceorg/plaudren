package plaud

import (
	"log/slog"
	"net/http"
)

type Server struct {
	server     *http.ServeMux
	listenAddr string
}

func New(listenAddr string) *Server {
	return &Server{
		server:     http.NewServeMux(),
		listenAddr: listenAddr,
	}
}

func (s *Server) Register(routers ...HTTPRouter) {
	for _, router := range routers {
		router.Register()
		router.RegisterServer(s.server)
	}
}

func (s *Server) Run() error {
	slog.Info("Server started on", "port", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, s.server)
}
