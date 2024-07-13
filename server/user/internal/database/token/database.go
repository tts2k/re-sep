package database

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"log/slog"
	"strings"
	"testing"
	"time"

	g "re-sep-user/internal/database/token/generated"
	config "re-sep-user/internal/system/config"

	"github.com/tursodatabase/go-libsql"
)

//go:embed schema/schema.sql
var schema string

var (
	dbPath  = config.Config().ConstructDBPath("token", "token.db")
	dbURL   = config.Config().TokenDB.URL
	token   = config.Config().TokenDB.Token
	db      *sql.DB
	queries *g.Queries
)

func InitTokenDB() {
	var dbCon *sql.DB
	var err error

	if testing.Testing() {
		dbCon, err = sql.Open("libsql", "file://"+dbPath)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		var connector *libsql.Connector
		connector, err = libsql.NewEmbeddedReplicaConnector(
			dbPath,
			dbURL,
			libsql.WithAuthToken(token),
		)
		if err != nil {
			log.Fatal(err)
		}
		dbCon = sql.OpenDB(connector)
	}

	db = dbCon
	queries = g.New(db)

	queryStrings := strings.Split(string(schema), ";\n")
	for _, query := range queryStrings {
		_, _ = db.Exec(query)
	}
}

func CloseTokenDB() {
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

func InsertToken(ctx context.Context, state string, userID string, duration time.Duration) *g.Token {
	expires := time.Now().Add(duration)
	params := g.InsertTokenParams{
		State:   state,
		Userid:  userID,
		Expires: expires,
	}

	result, err := queries.InsertToken(ctx, params)
	if err != nil {
		slog.Error("InsertToken:", "error", err)
		return nil
	}

	return &result
}

func RefreshToken(ctx context.Context, state string, duration time.Duration) *g.Token {
	expires := time.Now().Add(duration)
	params := g.RefreshTokenParams{
		State:   state,
		Expires: expires,
	}

	result, err := queries.RefreshToken(ctx, params)
	if err != nil {
		slog.Error("RefreshToken:", "error", err)
		return nil
	}

	return &result
}

func GetTokenByState(ctx context.Context, state string) *g.Token {
	result, err := queries.GetTokenByState(ctx, state)
	if err != nil {
		slog.Error("GetTokenByState:", "error", err)
		return nil
	}

	return &result
}

func DeleteToken(ctx context.Context, state string) g.Token {
	result, err := queries.DeleteToken(ctx, state)
	if err != nil {
		slog.Error("GetTokenByState:", "error", err)
		return result
	}

	return result
}
