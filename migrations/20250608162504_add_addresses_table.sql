-- +goose Up
-- +goose StatementBegin
CREATE TABLE addresses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    patient_id UUID NOT NULL,
    address_id UUID NOT NULL,
    house_number VARCHAR(5) DEFAULT NULL,
    house_body VARCHAR(5) DEFAULT NULL,
    house_apartment VARCHAR(5) DEFAULT NULL,
    fact BOOLEAN DEFAULT TRUE,
    active BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (patient_id) REFERENCES patients(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE addresses
-- +goose StatementEnd
