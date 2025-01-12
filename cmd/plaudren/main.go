package main

import (
	"log/slog"
	"plaudern/internal/api"
	"plaudern/internal/users"
)

func main() {
	server := api.New(":8000")

	userRouter := users.NewUserRouter()

	server.Register(userRouter)

	if err:= server.Run();err != nil {
		slog.Error(err.Error())
	}
}
