// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	State        string
	Token        string
	Expires      interface{}
	Refreshtoken string
}

type User struct {
	ID        uuid.UUID
	Name      string
	Sub       string
	LastLogin time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
