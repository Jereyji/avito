package main

import (
	"context"
	"fmt"
	"github.com/Jereyji/avito/internal/application/services"
	"github.com/Jereyji/avito/internal/infrastucture/database/postgres"
	"github.com/Jereyji/avito/internal/infrastucture/repository"
	"github.com/Jereyji/avito/internal/pkg/configs"
	"github.com/Jereyji/avito/internal/presentation/handlers"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	//if err := godotenv.Load("deployments/.env"); err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	cfgDatabase, err := configs.NewConfigDatabase()
	if err != nil {
		log.Fatal("Error reading environment variables: ", err)
	}

	//cfgDatabase.Host = "localhost"
	databaseConnectionStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfgDatabase.User, cfgDatabase.Password, cfgDatabase.Host, cfgDatabase.Port, cfgDatabase.Name, cfgDatabase.SSLMode,
	)

	postgresDB, err := postgres.NewPostgresDB(ctx, databaseConnectionStr)
	if err != nil {
		log.Fatal(err)
	}
	defer postgresDB.Close()

	var (
		repos   = repository.NewEstateRepository(postgresDB)
		service = services.NewEstateService(repos)
		handler = handlers.NewEstateHandler(service)
	)

	r := handler.InitRoutes()
	err = r.Run("0.0.0.0:8080")
	if err == nil {
		panic(err) // <?>
	}

}
