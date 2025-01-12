package api

import (
	"log/slog"
	"net/http"
)

/*
	TODO:
1. Middleware support
2. Ws support
3. Webrtc support
*/

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

func (s *ApiServer) Register(routers ...HTTPRouter) {
	for _, router := range routers {
		router.Register()
		router.RegisterServer(s.server)
	}
}

func (s *ApiServer) Run() error {
	slog.Info("Server started on", "port", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, s.server)
}
