-- name: GetTokenByState :one
SELECT * FROM Tokens
WHERE state = ? AND expires > Datetime('now')
LIMIT 1;

-- name: InsertToken :one
INSERT INTO Tokens (
	state, userId, expires
) VALUES (
	?, ?, ?
)
RETURNING *;

-- name: RefreshToken :one
UPDATE Tokens
SET expires = ?
WHERE state = ?
RETURNING *;

-- name: DeleteToken :one
DELETE FROM Tokens
WHERE state = ?
RETURNING *;

-- name: CleanTokens :exec
DELETE FROM Tokens
WHERE expires < Datetime('now');
