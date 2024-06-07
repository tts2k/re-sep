package database

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"log/slog"
	"strings"
	"time"

	g "re-sep-user/internal/database/generated"
	config "re-sep-user/internal/system/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema/schema.sql
var schema string

var (
	dbURL   = config.Config().DBURL
	db      *sql.DB
	queries *g.Queries
)

func InitDB() {
	dbCon, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	db = dbCon
	queries = g.New(db)

	queryStrings := strings.Split(string(schema), ";\n")
	for _, query := range queryStrings {
		db.Exec(query)
	}
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

func InsertToken(state string, userID string, duration time.Duration) *g.Token {
	expires := time.Now().Add(duration)
	params := g.InsertTokenParams{
		State:   state,
		Userid:  userID,
		Expires: &expires,
	}

	result, err := queries.InsertToken(context.Background(), params)
	if err != nil {
		slog.Error("InsertToken:", "error", err)
		return nil
	}

	return &result
}

func GetTokenByState(state string) g.Token {
	result, err := queries.GetTokenByState(context.Background(), state)
	if err != nil {
		slog.Error("GetTokenByState:", "error", err)
		return result
	}

	return result
}

func InsertUser(sub string, name string) *g.User {
	params := g.InsertUserParams{
		ID:   uuid.New(),
		Name: name,
		Sub:  sub,
	}

	user, err := queries.InsertUser(context.Background(), params)
	if err != nil {
		slog.Error("InsertUser:", "error", err)
		return nil
	}

	return &user
}

func GetUserByUniqueID(id string) *g.User {
	result, err := queries.GetUserByUniqueID(context.Background(), id)
	if err != nil {
		slog.Error("Cannot get token by unique ID", "error", err)
		return nil
	}

	return &result
}
