package main

import (
	"log"

	"github.com/dakstudios/auth-srv/db"
	"github.com/dakstudios/auth-srv/db/mongo"
	"github.com/dakstudios/auth-srv/handler"
	account "github.com/dakstudios/auth-srv/proto/account"
	auth "github.com/dakstudios/auth-srv/proto/auth"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("org.dakstudios.srv.auth"),
		micro.Flags(
			cli.StringFlag{
				Name:   "database_url",
				EnvVar: "DATABASE_URL",
				Usage:  "The Monogo database URL. Supports multiple hosts separated by comma",
			},
		),

		micro.Action(func(c *cli.Context) {
			if len(c.String("database_url")) > 0 {
				mongo.Url = c.String("database_url")
			}
		}),
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
