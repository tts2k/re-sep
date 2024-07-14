// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const cleanInactiveUsers = `-- name: CleanInactiveUsers :exec
DELETE FROM Users
WHERE CAST(Julianday(Datetime('now') - Julianday(last_login)) AS Integer) >= ?
`

func (q *Queries) CleanInactiveUsers(ctx context.Context, lastLogin string) error {
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
SELECT id, name, sub, config, last_login, created_at, updated_at FROM Users
WHERE sub = ? LIMIT 1
`

func (q *Queries) GetUserByUniqueID(ctx context.Context, sub string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUniqueID, sub)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Sub,
		&i.Config,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserConfig = `-- name: GetUserConfig :one
SELECT sub, config FROM v_user_config
WHERE sub = ? LIMIT 1
`

func (q *Queries) GetUserConfig(ctx context.Context, sub string) (VUserConfig, error) {
	row := q.db.QueryRowContext(ctx, getUserConfig, sub)
	var i VUserConfig
	err := row.Scan(&i.Sub, &i.Config)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO Users (
	id, name, sub, last_login, created_at, updated_at
) VALUES (
	?, ?, ?, Datetime('now'), Datetime('now'), Datetime('now')
)
RETURNING id, name, sub, config, last_login, created_at, updated_at
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
		&i.Config,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserConfig = `-- name: UpdateUserConfig :one
UPDATE v_user_config
SET config = ?
WHERE sub = ?
RETURNING sub, config
`

type UpdateUserConfigParams struct {
	Config string
	Sub    string
}

func (q *Queries) UpdateUserConfig(ctx context.Context, arg UpdateUserConfigParams) (VUserConfig, error) {
	row := q.db.QueryRowContext(ctx, updateUserConfig, arg.Config, arg.Sub)
	var i VUserConfig
	err := row.Scan(&i.Sub, &i.Config)
	return i, err
}

const updateUsername = `-- name: UpdateUsername :one
UPDATE Users
SET name = ?
WHERE sub = ?
RETURNING id, name, sub, config, last_login, created_at, updated_at
`

type UpdateUsernameParams struct {
	Name string
	Sub  string
}

func (q *Queries) UpdateUsername(ctx context.Context, arg UpdateUsernameParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUsername, arg.Name, arg.Sub)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Sub,
		&i.Config,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
