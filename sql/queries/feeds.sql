-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id) VALUES(uuid_generate_v4(), NOW(), NOW(), $1, $2, $3)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetFeeds :many
SELECT * FROM feeds WHERE user_id = $1 ORDER BY updated_at DESC;

-- name: GetFeed :one
SELECT * FROM feeds WHERE id = $1;