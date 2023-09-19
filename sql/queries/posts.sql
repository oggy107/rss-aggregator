-- name: CreatePost :one
INSERT INTO posts(id, created_at, updated_at, title, description, published_at, url, feed_id) VALUES(uuid_generate_v4(), NOW(), NOW(), $1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPostsForUsers :many
SELECT posts.* FROM posts JOIN feed_follows ON posts.feed_id = feed_follows.feed_id WHERE feed_follows.user_id = $1 ORDER BY posts.published_at DESC LIMIT $2;