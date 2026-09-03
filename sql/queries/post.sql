-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, published_at,title, url, description,feed_id)
VALUES (
           $1,
           $2,
           $3,
           $4,
           $5,
           $6,
        $7,
        $8
       )
    RETURNING *;

-- name: GetPostsForUser :many
select posts.*
from posts
    INNER JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
where feed_follows.user_id = $1
order by posts.published_at desc limit $2;




