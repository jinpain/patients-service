-- +goose Up
-- +goose StatementBegin
CREATE TABLE emails (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    patient_id UUID NOT NULL,
    email VARCHAR(32) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    work BOOLEAN NOT NULL DEFAULT TRUE,
    description VARCHAR(64) DEFAULT NULL,
    FOREIGN KEY (patient_id) REFERENCES patients(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE emails
-- +goose StatementEnd
