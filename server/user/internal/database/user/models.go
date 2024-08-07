// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Sub       string
	Config    sql.NullString
	LastLogin string
	CreatedAt string
	UpdatedAt string
}

type VUserConfig struct {
	Sub    string
	Config string
}
