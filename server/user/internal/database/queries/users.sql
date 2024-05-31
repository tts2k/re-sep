-- name: GetUserById :one
SELECT * FROM Users
WHERE id = ? LIMIT 1;

-- name: InsertUser :one
INSERT INTO Users (
	id, name, sub, last_login, created, updated
) VALUES (
	?, ?, ?, Datetime('now'), Datetime('now'), Datetime('now')
)
RETURNING *
;

-- name: UpdateUsername :one
UPDATE Users
SET name = ?
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = ?;

-- name: CleanInactiveUsers :exec
DELETE FROM Users
WHERE CAST(Julianday(Datetime('now') - Julianday(last_login)) AS Integer) >= ?;
