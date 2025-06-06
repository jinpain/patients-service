-- +goose Up
-- +goose StatementBegin
CREATE TABLE personal_documents(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    patient_id UUID NOT NULL,
    type SMALLINT DEFAULT 0,
    series CHAR(4) DEFAULT NULL,
    number CHAR(6) DEFAULT NULL,
    department_code CHAR(6) DEFAULT NULL,
    issue_date DATE DEFAULT NULL,
    company_id UUID DEFAULT NULL,
    nationality CHAR(2) DEFAULT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (patient_id) REFERENCES patients(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE personal_documents
-- +goose StatementEnd