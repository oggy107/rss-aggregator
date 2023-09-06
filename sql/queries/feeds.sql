-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id) VALUES(uuid_generate_v4(), NOW(), NOW(), $1, $2, $3)
RETURNING *;