-- name: CreateUser :exec
INSERT INTO user (
  username, password, role
) VALUES (
  ?, ?, ?
);

-- name: GetUserByID :one
SELECT * FROM user WHERE id = ?;

-- name: GetUserByUsernameAndPassword :one
SELECT * FROM user WHERE username = ? AND password = ?;

-- name: GetAdminUser :one
SELECT * FROM user WHERE username = "admin";