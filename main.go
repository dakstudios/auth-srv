package main

import (
	"log"

	"github.com/dakstudios/auth-srv/db"
	_ "github.com/dakstudios/auth-srv/db/mongo"
	"github.com/dakstudios/auth-srv/handler"
	account "github.com/dakstudios/auth-srv/proto/account"
	auth "github.com/dakstudios/auth-srv/proto/auth"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("org.dakstudios.srv.auth"),
	)

	service.Init()

	account.RegisterAccountServiceHandler(service.Server(), new(handler.Account))
	auth.RegisterAuthHandler(service.Server(), new(handler.Auth))

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
