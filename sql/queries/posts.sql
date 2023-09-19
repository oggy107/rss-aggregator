-- name: CreatePost :one
INSERT INTO posts(id, created_at, updated_at, title, description, published_at, url, feed_id) VALUES(uuid_generate_v4(), NOW(), NOW(), $1, $2, $3, $4, $5)
RETURNING *;