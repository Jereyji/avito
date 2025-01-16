package configs

import "github.com/ilyakaznacheev/cleanenv"

type ConfigDatabase struct {
	Host     string `env:"POSTGRES_HOST" env-default:"postgres"`
	Port     string `env:"POSTGRES_PORT" env-default:"5432"`
	Name     string `env:"POSTGRES_DB" env-default:"avito_db"`
	User     string `env:"POSTGRES_USER" env-default:"user"`
	Password string `env:"POSTGRES_PASSWORD"`
	SSLMode  string `env:"POSTGRES_SSLMODE" env-default:"disable"`
}

func NewConfigDatabase() (*ConfigDatabase, error) {
	var cfgDatabase ConfigDatabase

	err := cleanenv.ReadEnv(&cfgDatabase)
	if err != nil {
		return nil, err
	}

	return &cfgDatabase, nil
}
