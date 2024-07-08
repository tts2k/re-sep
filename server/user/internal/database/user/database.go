package database

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"strings"
	"time"

	g "re-sep-user/internal/database/user/generated"
	config "re-sep-user/internal/system/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema/schema.sql
var schema string

var (
	dbURL   = config.Config().ConstructDBPath("user.db")
	db      *sql.DB
	queries *g.Queries
)

var DefaultUserConfig = UserConfig{
	Font:     "serif",
	FontSize: 3,
	Justify:  false,
	Margin: Margin{
		Left:  3,
		Right: 3,
	},
}

func InitUserDB() {
	dbCon, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	db = dbCon
	queries = g.New(db)

	splitted := strings.Split(string(schema), ";\n")
	var queries []string

	for _, split := range splitted {

		lowered := strings.ToLower(split)
		if strings.Contains(lowered, "end") {
			queries[len(queries)-1] = queries[len(queries)-1] + " ;END;"
			continue
		}

		queries = append(queries, split)
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Println(query)
			log.Fatal(err)
		}
	}
}

func CloseUserDB() {
	db.Close()
}

func Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func InsertUser(ctx context.Context, sub string, name string) *g.User {
	params := g.InsertUserParams{
		ID:   uuid.New(),
		Name: name,
		Sub:  sub,
	}

	user, err := queries.InsertUser(ctx, params)
	if err != nil {
		slog.Error("Cannot insert user", "database_error", err)
		return nil
	}

	return &user
}

func GetUserByUniqueID(ctx context.Context, id string) *g.User {
	result, err := queries.GetUserByUniqueID(ctx, id)
	if err != nil {
		slog.Error("Cannot get user by unique ID", "database_error", err)
		return nil
	}

	return &result
}

func UpdateUsername(ctx context.Context, sub string, username string) *g.User {
	params := g.UpdateUsernameParams{
		Name: username,
		Sub:  sub,
	}

	result, err := queries.UpdateUsername(ctx, params)
	if err != nil {
		slog.Error("Cannot update username", "database_error", err)
		return nil
	}

	return &result
}

func UpdateUserConfig(ctx context.Context, sub string, config *UserConfig) *UserConfig {
	jsonConfig, err := json.Marshal(config)
	if err != nil {
		slog.Error("Cannot update username", "json_error", err)
		return nil
	}

	params := g.UpdateUserConfigParams{
		Sub:    sub,
		Config: string(jsonConfig),
	}

	uc, err := queries.UpdateUserConfig(ctx, params)
	if err != nil {
		slog.Error("Cannot update user config", "database_error", err)
		return nil
	}

	var result UserConfig
	err = json.Unmarshal([]byte(uc.Config), &result)
	if err != nil {
		slog.Error("Cannot parse user config", "json_error", err)
	}

	return &result
}

func GetUserConfig(ctx context.Context, sub string) *UserConfig {
	result, err := queries.GetUserConfig(ctx, sub)
	if err != nil {
		slog.Error("Cannot get user config", "database_error", err)
		return nil
	}

	var userConfig UserConfig
	json.Unmarshal([]byte(result.Config), &userConfig)

	return &userConfig
}
