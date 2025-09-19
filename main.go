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
	"github.com/redis/go-redis/v9"
)

// @title SocialNetwork API
// @contact.name SocialNetwork API Service
// @contact.url https://socialnetwork.com
// @contact.email help@socialetwork.com
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

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",
			configs.AppSettings.RedisParams.Host,
			configs.AppSettings.RedisParams.Port),
		DB: configs.AppSettings.RedisParams.Database,
	})

	cache := repository.NewCache(rdb)

	repository := repository.NewRepository(db, cache)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	if err = controller.RunServer(fmt.Sprintf(":%s", configs.AppSettings.AppParams.PortRun)); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
