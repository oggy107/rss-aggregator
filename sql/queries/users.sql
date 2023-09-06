-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name, api_key) VALUES(uuid_generate_v4(), NOW(), NOW(), $1, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;
