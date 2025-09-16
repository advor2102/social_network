package main

import (
	"log"

	"github.com/advor2102/socialnetwork/internal/controller"
	"github.com/advor2102/socialnetwork/internal/repository"
	"github.com/advor2102/socialnetwork/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dns := "host=localhost port=5432 user=postgres password=!Makar24052018 dname=social_network_db sslmode=disable"

	db, err := sqlx.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	if err = controller.RunServer(":7777"); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
