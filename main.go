package main

import (
	"fmt"
	"log"
	"os"

	"github.com/advor2102/socialnetwork/internal/configs"
	"github.com/advor2102/socialnetwork/internal/controller"
	"github.com/advor2102/socialnetwork/internal/repository"
	"github.com/advor2102/socialnetwork/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	if err := configs.ReadSettings(); err != nil {
		log.Fatal(err)
	}

	dns := fmt.Sprintf(`host=%s 
						port=%s
						user=%s 
						password=%s 
						dbname=%s 
						sslmode=disable`,
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		os.Getenv("POSTGRES_PASSWORD"),
		configs.AppSettings.PostgresParams.Database,
	)

	db, err := sqlx.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	if err = controller.RunServer(fmt.Sprintf(":%s", configs.AppSettings.AppParams.PortRun)); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
