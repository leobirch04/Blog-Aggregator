-- +goose Up
CREATE TABLE users
(
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR NOT NULL
);

-- +goose Down
DROP TABLE users;