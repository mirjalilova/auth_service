package main

import (
	"auth-service/api"
	"auth-service/api/handlers"
	"auth-service/config"
	"auth-service/postgresql"
	"auth-service/service"
)

func main() {
	cf := config.Load()

	conn, err := postgresql.ConnectDB(&cf)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	us := service.NewUserService(conn)
	handler := handlers.NewHandler(us)

	roter := api.NewRouter(handler)
	if err := roter.Run(cf.AUTH_PORT); err != nil {
		panic(err)
	}
}
