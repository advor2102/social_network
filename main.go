package main

import (
	"fmt"
	"os"

	"github.com/advor2102/socialnetwork/internal/configs"
	"github.com/advor2102/socialnetwork/internal/controller"
	"github.com/advor2102/socialnetwork/internal/repository"
	"github.com/advor2102/socialnetwork/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

// @title SocialNetwork API
// @contact.name SocialNetwork API Service
// @contact.url https://socialnetwork.com
// @contact.email help@socialetwork.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Info().Msg("Starting up application...")

	if err := configs.ReadSettings(); err != nil {
		logger.Error().Err(err).Msg("Error during reading settings")
		return
	}
	logger.Info().Msg("Read settings successfully")

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
		logger.Error().Err(err).Msg("Error during connection to database")
	}
	logger.Info().Msg("Database connected successfully")

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",
			configs.AppSettings.RedisParams.Host,
			configs.AppSettings.RedisParams.Port),
		DB: configs.AppSettings.RedisParams.Database,
	})

	cache := repository.NewCache(rdb)
	logger.Info().Msg("Redis connected successfully")

	repository := repository.NewRepository(db)
	service := service.NewService(repository, cache)
	controller := controller.NewController(service)

	if err = controller.RunServer(fmt.Sprintf(":%s", configs.AppSettings.AppParams.PortRun)); err != nil {
		logger.Error().Err(err).Msg("Error during running http server")
	}

	if err = db.Close(); err != nil {
		logger.Error().Err(err).Msg("Error during closing database connection")
	}
}
