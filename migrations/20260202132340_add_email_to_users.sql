-- +goose Up
ALTER TABLE users ADD COLUMN email VARCHAR(100) NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE users DROP COLUMN email;
