-- name: CreateFeedFollows :one
INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id) VALUES(uuid_generate_v4(), NOW(), NOW(), $1, $2)
RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE user_id = $1 ORDER BY updated_at DESC;
