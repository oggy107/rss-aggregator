-- name: CreateFeedFollows :one
INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id) VALUES(uuid_generate_v4(), NOW(), NOW(), $1, $2)
RETURNING *;