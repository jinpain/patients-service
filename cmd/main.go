package main

import (
	"github.com/his-vita/patients-service/internal/config"
	"github.com/his-vita/patients-service/pkg/database/postgres"
	"github.com/his-vita/patients-service/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	pgContext, err := postgres.NewPostgresConnect(&cfg.Db)
	if err != nil {
		panic(err)
	}

	postgres.NewTransactionManager(pgContext)

	log.Info("application running")
}
