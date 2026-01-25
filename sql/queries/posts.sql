-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT f.name, p.title, p.description, p.url, p.published_at, p.id, p.feed_id
FROM posts AS p
JOIN feed_follows AS ff
ON p.feed_id = ff.feed_id
JOIN feeds AS f
ON ff.feed_id = f.id
JOIN users AS u
ON ff.user_id = u.id
WHERE u.id = $1
ORDER BY p.published_at DESC
LIMIT $2;