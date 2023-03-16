package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/vaberof/it-planet-task/internal/app/http/handler"
	"github.com/vaberof/it-planet-task/internal/domain/account"
	"github.com/vaberof/it-planet-task/internal/domain/animal"
	"github.com/vaberof/it-planet-task/internal/domain/animaltype"
	"github.com/vaberof/it-planet-task/internal/domain/location"
	"github.com/vaberof/it-planet-task/internal/domain/vstlocation"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/accountpg"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/animalpg"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/animaltypepg"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/locationpg"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/vstlocationpg"
	"github.com/vaberof/it-planet-task/internal/service/auth"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("cannot initialize config: %s", err.Error())
	}

	if err := loadEnvironmentVariables(); err != nil {
		log.Fatalf("cannot load environment variables: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(&postgres.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Name:     viper.GetString("database.name"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("cannot connect to database %s", err.Error())
	}

	err = db.AutoMigrate(&accountpg.Account{}, &locationpg.Location{}, &animalpg.Animal{}, &animaltypepg.AnimalType{}, &vstlocationpg.VisitedLocation{})
	if err != nil {
		log.Fatalf("cannot auto migrate models %s", err.Error())
	}

	accountStorage := accountpg.NewPostgresAccountStorage(db)
	locationStorage := locationpg.NewPostgresLocationStorage(db)
	animalTypeStorage := animaltypepg.NewPostgresAnimalTypeStorage(db)
	visitedLocationStorage := vstlocationpg.NewPostgresVisitedLocationStorage(db)
	animalStorage := animalpg.NewPostgresAnimalStorage(db, visitedLocationStorage)

	accountService := account.NewAccountService(accountStorage)
	locationService := location.NewLocationService(locationStorage)
	visitedLocationService := vstlocation.NewVisitedLocationService(visitedLocationStorage, locationStorage)
	animalTypeService := animaltype.NewAnimalTypeService(animalTypeStorage)
	animalService := animal.NewAnimalService(animalStorage, animalTypeStorage, accountStorage, locationStorage, visitedLocationService)

	authService := auth.NewAuthService(accountStorage)

	httpHandler := handler.NewHttpHandler(accountService, locationService, animalTypeService, animalService, authService)

	router := httpHandler.InitRouter()

	if err = router.Run(viper.GetString("server.host") + ":" + viper.GetString("server.port")); err != nil {
		log.Fatalf("cannot run a server: %v", err)
	}
}

func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func loadEnvironmentVariables() error {
	err := godotenv.Load("./.env")
	return err
}
