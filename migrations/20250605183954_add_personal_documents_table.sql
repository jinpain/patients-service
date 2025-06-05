-- +goose Up
-- +goose StatementBegin
CREATE TABLE personal_documents(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    patient_id UUID NOT NULL,
    series CHAR(4) DEFAULT NULL,
    number CHAR(6) DEFAULT NULL,
    department_code CHAR(6) DEFAULT NULL,
    issue_date DATE DEFAULT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (patient_id) REFERENCES patients(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE personal_documents
-- +goose StatementEnd