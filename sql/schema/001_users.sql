-- +goose Up
CREATE TABLE if not exists users
(
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR NOT NULL unique
);

create table if not exists feeds
(
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR NOT NULL,
    url varchar not null unique,
    user_id UUID NOT NULL references users(id) on delete cascade,
    last_fetched_at TIMESTAMP
);

create table if not exists feed_follows
(
    id UUID primary key NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id UUID NOT NULL references users(id) on delete cascade,
    feed_id UUID not null references feeds(id) on delete cascade,
    UNIQUE (user_id, feed_id)
);

create table if not exists posts
(
    id UUID primary key NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    published_at TIMESTAMP NOT NULL DEFAULT NOW(),
    title varchar,
    url varchar not null unique,
    description varchar,
    feed_id UUID not null references feeds(id) on delete cascade
);


-- +goose Down
drop table feed_follows;
drop table feeds;
DROP TABLE users;
DROP TABLE posts;
