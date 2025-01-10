package users

import (
	"net/http"
	"plaudern/internal/api"
)

type UserRouter struct {
	*api.Router
}

func NewUserRouter() *UserRouter {
	return &UserRouter{
		Router: api.NewRouter("/user"),
	}
}

func (s *UserRouter) Register() {
	s.Router.Get("/login", s.Login)
}

func (s *UserRouter) Login(w http.ResponseWriter, r *http.Request) *api.ApiError {
	w.Write([]byte("ok\r\n"))
	return nil
}
