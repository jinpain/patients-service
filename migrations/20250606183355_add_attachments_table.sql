-- +goose Up
-- +goose StatementBegin
CREATE TABLE attachments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    patient_id UUID NOT NULL,
    lpu_id UUID NOT NULL,
    department_id UUID NOT NULL,
    purpose SMALLINT NOT NULL,
    type SMALLINT NOT NULL,
    sector_code SMALLINT NOT NULL,
    reg_date DATE NOT NULL,
    cause SMALLINT DEFAULT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (patient_id) REFERENCES patients(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE attachments
-- +goose StatementEnd
