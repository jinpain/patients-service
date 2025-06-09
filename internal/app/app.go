package app

import (
	"github.com/his-vita/patients-service/internal/config"
	"github.com/his-vita/patients-service/internal/repository"
	"github.com/his-vita/patients-service/pkg/database/postgres"
	"github.com/his-vita/patients-service/pkg/logger"
	sqlstore "github.com/his-vita/patients-service/pkg/sql-store"
)

func Run(cfg *config.Config) {
	log := logger.New(cfg.Env)

	pgContext, err := postgres.NewPostgresConnect(&cfg.Db)
	if err != nil {
		panic(err)
	}

	sqlStore, err := sqlstore.New(cfg.Sql.Path)
	if err != nil {
		panic(err)
	}

	postgres.NewTransactionManager(pgContext)

	addressRepo := repository.NewAddressRepository(pgContext, sqlStore)
	anthropometryRepo := repository.NewAnthropometryRepository(pgContext, sqlStore)
	attahmentRepo := repository.NewAttahmentRepository(pgContext, sqlStore)
	emailRepo := repository.NewEmailRepository(pgContext, sqlStore)
	insurancePolicyRepo := repository.NewInsurancePolicyRepository(pgContext, sqlStore)
	patientRepo := repository.NewPatientRepository(pgContext, sqlStore)
	personalDocumentRepo := repository.NewPersonalDocumentRepository(pgContext, sqlStore)
	phoneRepo := repository.NewPhoneRepository(pgContext, sqlStore)

	log.Info("application running")
}
