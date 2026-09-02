-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
           $1,
           $2,
           $3,
           $4,
           $5,
           $6
       )
    RETURNING *;

-- name: GetAllFeeds :many
select * from feeds;

-- name: GetFeedByURL :one
select * from feeds where url = $1 LIMIT 1;



-- name: ClearFeeds :exec
delete from feeds;