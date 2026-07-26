-- name: CreateUser :one
INSERT INTO users (email, name)
VALUES ($1, $2)
RETURNING id, email, name;

-- name: GetUser :one
SELECT id, email, name
FROM users
WHERE id = $1;
