-- +goose Up
-- +goose StatementBegin
CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100) DEFAULT NULL,
    birth_date DATE NOT NULL,
    gender bool DEFAULT NULL,
    snils CHAR(11) DEFAULT NULL,
    inn CHAR(12) DEFAULT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE patients
-- +goose StatementEnd
