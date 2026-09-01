-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
           $1,
           $2,
           $3,
           $4
       )
    RETURNING *;

-- name: GetUser :one
select * from users where name = $1 LIMIT 1;

-- name: GetUserName :one
select name from users where id = $1 LIMIT 1;

-- name: UserExists :one
select EXISTS(select 1 from users  where name = $1) ;

-- name: ClearUsers :exec
delete from users;

-- name: GetAllUsers :many
select * from users;