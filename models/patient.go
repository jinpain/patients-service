package models

import (
	"cloud.google.com/go/civil"
	"github.com/google/uuid"
)

type Patient struct {
	Id          uuid.UUID  `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	MiddleName  string     `json:"middle_name"`
	BirthDate   civil.Date `json:"birth_date"`
	PhoneNumber string     `json:"phone_number"`
	Email       string     `json:"email"`
	Version     int        `json:"version"`
}
