-- +goose Up
-- +goose StatementBegin
CREATE TABLE insurance_policies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    patient_id UUID NOT NULL,
    number CHAR(16) NOT NULL,
    issue_date DATE NOT NULL,
    type SMALLINT DEFAULT NULL,
    kind SMALLINT DEFAULT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    company_id UUID DEFAULT NULL,
    FOREIGN KEY (patient_id) REFERENCES patients(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE insurance_policies
-- +goose StatementEnd