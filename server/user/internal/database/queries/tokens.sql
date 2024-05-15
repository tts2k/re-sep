-- name: GetTokenById :one
SELECT * FROM Tokens
WHERE id = ? LIMIT 1;

-- name: GetUserByTokenId :one
SELECT * FROM Tokens
WHERE id = ? LIMIT 1;

-- name: InsertToken :one
INSERT INTO Tokens (
	id, userId, expires
) VALUES (
	?, ?, ?
)
RETURNING *
;

-- name: UpdateToken :one
UPDATE Tokens
SET expires = ?
WHERE id = ?
RETURNING *;

-- name: CleanTokens :exec
DELETE FROM Tokens
WHERE expires < Datetime("now")
