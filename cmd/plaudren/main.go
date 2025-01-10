package main

import (
	"flag"
	"log/slog"
	"plaudern/internal/api"
	"plaudern/internal/users"
)

func main() {
	var (
		listenAddr = flag.String("host", ":8000", "--host :<port> | hostAddr ")
	)
	flag.Parse()
	server := api.New(*listenAddr)
	userRouter := users.NewUserRouter()
	server.Register(userRouter)
	err:=server.Run()
	if err!= nil {
		slog.Error(err.Error())
	}
}
