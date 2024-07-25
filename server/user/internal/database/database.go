package database

import (
	_ "embed"

	config "re-sep-user/internal/system/config"
)

//go:embed user/schema/schema.sql
var userSchema string

//go:embed token/schema/schema.sql
var tokenSchema string

var systemConfig = config.Config()

type DB interface {
	Migrate()
}
