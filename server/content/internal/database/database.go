package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tursodatabase/go-libsql"
)

type Service interface {
	Health() map[string]string
	GetArticle(context context.Context, entryName string) (*Article, error)
}

type service struct {
	db *sql.DB
}

var (
	dbPath     = os.Getenv("DB_PATH")
	dbURL      = os.Getenv("DB_URL")
	authToken  = os.Getenv("DB_AUTH_TOKEN")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	var db *sql.DB
	var err error
	// if turso
	var connector *libsql.Connector
	connector, err = libsql.NewEmbeddedReplicaConnector(dbPath, dbURL,
		libsql.WithAuthToken(authToken),
	)
	db = sql.OpenDB(connector)
	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func (s *service) GetArticle(ctx context.Context, entryName string) (*Article, error) {
	row := s.db.QueryRow(`
		SELECT title, entry_name, issued, modified, html_text, json(author), json(toc)
		FROM Articles WHERE entry_name = ?
	`, entryName)
	if row.Err() != nil {
		slog.Error(row.Err().Error())
		return nil, nil
	}

	article := Article{}
	err := row.Scan(
		&article.Title,
		&article.EntryName,
		&article.Issued,
		&article.Modified,
		&article.HTMLText,
		&article.Author,
		&article.TOC,
	)

	return &article, err
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
