package repository

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/database"
	"github.com/his-vita/patients-service/models"
	"github.com/his-vita/patients-service/pkg/sqlutils"
	"github.com/jackc/pgx"
)

type PatientRepository struct {
	log       *slog.Logger
	pgContext *database.PgContext
	sqlFiles  map[string]string
}

func NewPatientRepository(log *slog.Logger, pgContext *database.PgContext, sqlPath string) *PatientRepository {
	filePath := filepath.Join(sqlPath, "patients")
	if err := sqlutils.CheckSQLFilesPath(filePath); err != nil {
		panic(err)
	}

	sqlFiles, err := sqlutils.LoadSQLFiles(filePath)
	if err != nil {
		panic(err)
	}

	return &PatientRepository{
		log:       log,
		pgContext: pgContext,
		sqlFiles:  sqlFiles,
	}
}

func (pr *PatientRepository) GetPatient(id *uuid.UUID) (*models.Patient, error) {
	query, exists := pr.sqlFiles["get_patient_by_id.sql"]
	if !exists {
		return nil, fmt.Errorf("SQL query insert_patient not found")
	}

	var patient models.Patient

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	err := pr.pgContext.Pool.QueryRow(ctx, query, id).Scan(
		&patient.Id,
		&patient.FirstName,
		&patient.LastName,
		&patient.MiddleName,
		&patient.BirthDate,
		&patient.PhoneNumber,
		&patient.Email,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("patient with id %s not found", id)
		}
		return nil, fmt.Errorf("error retrieving patient: %w", err)
	}

	return &patient, nil
}

func (pr *PatientRepository) GetPatients(limit int, offset int) ([]models.Patient, error) {
	query, exists := pr.sqlFiles["get_patients.sql"]
	if !exists {
		return nil, fmt.Errorf("SQL query get_patients not found")
	}

	var patients []models.Patient

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	rows, err := pr.pgContext.Pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var patient models.Patient
		if err := rows.Scan(
			&patient.Id,
			&patient.FirstName,
			&patient.LastName,
			&patient.MiddleName,
			&patient.BirthDate,
			&patient.PhoneNumber,
			&patient.Email,
			&patient.Version); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		patients = append(patients, patient)
	}

	return patients, nil
}

func (pr *PatientRepository) UpdatePatient(patient *models.Patient) error {
	query, exists := pr.sqlFiles["update_patient.sql"]
	if !exists {
		return fmt.Errorf("SQL query update_patient not found")
	}

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	res, err := pr.pgContext.Pool.Exec(ctx, query, patient.Id, patient.FirstName, patient.LastName, patient.MiddleName, patient.BirthDate, patient.PhoneNumber, patient.Email, "admin", patient.Version)
	if err != nil {
		return fmt.Errorf("error update patient: %w", err)
	}

	rowsAffected := res.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("error update patient. Version in DB is higher")
	}

	return nil
}

func (pr *PatientRepository) CreatePatient(patient *models.Patient) error {
	query, exists := pr.sqlFiles["insert_patient.sql"]
	if !exists {
		return fmt.Errorf("SQL query insert_patient not found")
	}

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err := pr.pgContext.Pool.Exec(ctx, query, patient.FirstName, patient.LastName, patient.MiddleName, patient.BirthDate, patient.PhoneNumber, patient.Email, "admin")
	if err != nil {
		return fmt.Errorf("error creating patient: %w", err)
	}

	return nil
}

func (pr *PatientRepository) MarkPatientAsDeleted(id *uuid.UUID) error {
	query, exists := pr.sqlFiles["mark_deleted_patient.sql"]
	if !exists {
		return fmt.Errorf("SQL query mark_deleted_patient not found")
	}

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err := pr.pgContext.Pool.Exec(ctx, query, id, "admin")
	if err != nil {
		return fmt.Errorf("error mark deleted patient: %w", err)
	}

	return nil
}

func (pr *PatientRepository) UnMarkPatientAsDeleted(id *uuid.UUID) error {
	query, exists := pr.sqlFiles["unmark_deleted_patient.sql"]
	if !exists {
		return fmt.Errorf("SQL query unmark_deleted_patient not found")
	}

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err := pr.pgContext.Pool.Exec(ctx, query, id, "admin")
	if err != nil {
		return fmt.Errorf("error unmark deleted patient: %w", err)
	}

	return nil
}
