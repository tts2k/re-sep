package database

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	"re-sep-scraper/scraper"
)

var dbPath string

func InitDB(path string) {
	dbPath = path
}

func connectDB() (*sql.DB, error) {
	return sql.Open("sqlite3", dbPath)
}

func CreateTable() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS articles (
			id uuid PRIMARY KEY,
			title text NOT NULL,
			entry_name text NOT NULL UNIQUE,
			issued timestamp,
			modified timestamp,
			html_text text,
			author blob,
			toc blob,

			UNIQUE(entry_name)
		)`,
	)
	if err != nil {
		return err
	}

	return nil
}

func InsertArticle(article scraper.Article) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(`
		INSERT INTO articles (
			id, title, entry_name, issued, modified, html_text, author, toc
		) VALUES (
			?, ?, ?, ?, ?, ?, jsonb(?), jsonb(?)
		)
		ON CONFLICT (entry_name) DO UPDATE SET
			title=excluded.title,
			issued=excluded.issued,
			modified=excluded.modified,
			html_text=excluded.html_text,
			author=excluded.author,
			author=excluded.toc
	`)
	if err != nil {
		return err
	}

	// Format date
	// https://www.sqlite.org/lang_datefunc.html
	issuedDate := article.Issued.Format(time.RFC3339)
	modifiedDate := article.Modified.Format(time.RFC3339)

	// Create JSON objects
	// https://www.sqlite.org/json1.html
	authorsJSON, err := json.Marshal(article.Author)
	if err != nil {
		return err
	}

	tocJSON, err := json.Marshal(article.TOC)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		uuid.NewString(),
		article.Title,
		article.EntryName,
		issuedDate,
		modifiedDate,
		article.HTMLText,
		string(authorsJSON),
		string(tocJSON),
	)
	if err != nil {
		return err
	}

	return nil
}
