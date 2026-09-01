-- +goose Up
CREATE TABLE users
(
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR NOT NULL unique
);

create table feeds
(
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR NOT NULL,
    url varchar not null unique,
    user_id UUID NOT NULL references users(id) on delete cascade
);

-- +goose Down
drop table feeds;
DROP TABLE users;