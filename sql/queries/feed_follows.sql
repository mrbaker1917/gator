-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (user_id, feed_id)
    VALUES ($1, $2)
    RETURNING *
)
SELECT 
    inserted_feed_follow.id, 
    inserted_feed_follow.created_at, 
    inserted_feed_follow.updated_at, 
    inserted_feed_follow.user_id, 
    inserted_feed_follow.feed_id, 
    u.name AS user_name, 
    f.name AS feed_name
FROM inserted_feed_follow
JOIN users AS u 
    ON u.id = inserted_feed_follow.user_id
JOIN feeds AS f 
    ON f.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT 
    ff.id,
    ff.created_at,
    ff.updated_at,
    ff.user_id,
    ff.feed_id,
    u.name AS user_name, 
    f.name AS feed_name
FROM feed_follows as ff 
JOIN users AS u
    ON u.id = ff.user_id
JOIN feeds AS f
    ON f.id = ff.feed_id
WHERE u.id = $1;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows
WHERE user_id = $1 
AND feed_id = $2; 