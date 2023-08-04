-- name: CreateUser :one
INSERT INTO users (
  email, 
  password
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM users
ORDER BY email;

-- name: UpdateUser :exec
UPDATE users 
SET password = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;