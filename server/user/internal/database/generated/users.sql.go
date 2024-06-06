// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const cleanInactiveUsers = `-- name: CleanInactiveUsers :exec
DELETE FROM Users
WHERE CAST(Julianday(Datetime('now') - Julianday(last_login)) AS Integer) >= ?
`

func (q *Queries) CleanInactiveUsers(ctx context.Context, lastLogin time.Time) error {
	_, err := q.db.ExecContext(ctx, cleanInactiveUsers, lastLogin)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByUniqueID = `-- name: GetUserByUniqueID :one
SELECT id, name, sub, last_login, created_at, updated_at FROM Users
WHERE sub = ? LIMIT 1
`

func (q *Queries) GetUserByUniqueID(ctx context.Context, sub string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUniqueID, sub)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Sub,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO Users (
	id, name, sub, last_login, created_at, updated_at
) VALUES (
	?, ?, ?, Datetime('now'), Datetime('now'), Datetime('now')
)
RETURNING id, name, sub, last_login, created_at, updated_at
`

type InsertUserParams struct {
	ID   uuid.UUID
	Name string
	Sub  string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser, arg.ID, arg.Name, arg.Sub)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Sub,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUsername = `-- name: UpdateUsername :one
;

UPDATE Users
SET name = ?
WHERE id = ?
RETURNING id, name, sub, last_login, created_at, updated_at
`

type UpdateUsernameParams struct {
	Name string
	ID   uuid.UUID
}

func (q *Queries) UpdateUsername(ctx context.Context, arg UpdateUsernameParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUsername, arg.Name, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Sub,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
