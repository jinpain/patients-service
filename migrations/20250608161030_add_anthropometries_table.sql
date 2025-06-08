-- +goose Up
-- +goose StatementBegin
CREATE TABLE anthropometries (
    patient_id UUID PRIMARY KEY,
    blood_group SMALLINT DEFAULT NULL,
    rh_factor SMALLINT DEFAULT NULL,
    eye_color SMALLINT DEFAULT NULL,
    kell_affiliation SMALLINT DEFAULT NULL,
    height float DEFAULT NULL,
    weight float DEFAULT NULL,
    FOREIGN KEY (patient_id) REFERENCES patients(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE anthropometries
-- +goose StatementEnd
