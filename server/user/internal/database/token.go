package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tursodatabase/go-libsql"

	g "re-sep-user/internal/database/token"
)

type TokenDB struct {
	conn    *sql.DB
	Queries *g.Queries
}

func NewTokenDB() *TokenDB {
	connector, err := libsql.NewEmbeddedReplicaConnector(
		systemConfig.TokenDB.Path,
		systemConfig.TokenDB.URL,
		libsql.WithAuthToken(systemConfig.TokenDB.Token),
	)
	if err != nil {
		log.Fatal(err)
	}
	dbCon := sql.OpenDB(connector)
	queries := g.New(dbCon)

	return &TokenDB{
		conn:    dbCon,
		Queries: queries,
	}
}

func NewTokenDBMemory() *TokenDB {
	dbCon, err := sql.Open("libsql", "file::memory:?cache=shared&mode=rwc&_journal_mode=WAL&busy_timeout=10000")
	dbCon.SetMaxOpenConns(1)
	if err != nil {
		log.Fatal(err)

	}

	queries := g.New(dbCon)

	return &TokenDB{
		conn:    dbCon,
		Queries: queries,
	}
}

func (db *TokenDB) Migrate() {
	splitted := strings.Split(string(tokenSchema), ";\n")
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
		if query == "" {
			continue
		}

		_, err := db.conn.Exec(query)
		if err != nil {
			log.Println(query)
			log.Fatal(err)
		}
	}
}

func (db *TokenDB) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := db.conn.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
