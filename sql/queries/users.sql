-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name) VALUES(uuid_generate_v4(), NOW(), NOW(), $1)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;
