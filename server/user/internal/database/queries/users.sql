-- name: GetUserById :one
SELECT * FROM Users
WHERE id = ? LIMIT 1;

-- name: InsertUser :one
INSERT INTO Users (
	id, username, password, email, created, updated
) VALUES (
	?, ?, ?, ?, Datetime('now'), Datetime('now')
)
RETURNING *
;

-- name: UpdateUserPassword :one
UPDATE Users
SET password = ?
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = ?
