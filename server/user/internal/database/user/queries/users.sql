-- name: GetUserByUniqueID :one
SELECT * FROM Users
WHERE sub = ? LIMIT 1;

-- name: InsertUser :one
INSERT INTO Users (
	id, name, sub, last_login, created_at, updated_at
) VALUES (
	?, ?, ?, Datetime('now'), Datetime('now'), Datetime('now')
)
RETURNING *;

-- name: GetUserConfig :one
SELECT * FROM v_user_config
WHERE sub = ? LIMIT 1;

-- name: UpdateUsername :one
UPDATE Users
SET name = ?
WHERE sub = ?
RETURNING *;

-- name: UpdateUserConfig :one
UPDATE v_user_config
SET config = ?
WHERE sub = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = ?;

-- name: CleanInactiveUsers :exec
DELETE FROM Users
WHERE CAST(Julianday(Datetime('now') - Julianday(last_login)) AS Integer) >= ?;
