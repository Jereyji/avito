package main

import (
	"context"
	"fmt"
	"github.com/Jereyji/avito/internal/application/services"
	"github.com/Jereyji/avito/internal/infrastucture/postgres"
	"github.com/Jereyji/avito/internal/infrastucture/repository"
	"github.com/Jereyji/avito/internal/presentation/handlers"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os/signal"
	"syscall"
)

type ConfigDatabase struct {
	Port     string `env:"DB_PORT" env-default:"5432"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Name     string `env:"DB_NAME" env-default:"postgres"`
	User     string `env:"DB_USERNAME" env-default:"user"`
	Password string `env:"DB_PASSWORD"`
	SSLMode  string `env:"DB_SSLMODE" env-default:"disable"`
}

func main() {
	// Load Config

	if err := godotenv.Load("deployments/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfgDatabase ConfigDatabase
	if err := cleanenv.ReadEnv(&cfgDatabase); err != nil {
		log.Fatal("Error reading environment variables: ", err)
	}

	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfgDatabase.User, cfgDatabase.Password, cfgDatabase.Host, cfgDatabase.Port, cfgDatabase.Name, cfgDatabase.SSLMode,
	)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	postgresDB, err := postgres.NewPostgresDB(ctx, connString)
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
	err = r.Run()
	if err == nil {
		panic(err) // <?>
	}
}
