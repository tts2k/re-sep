package database

import (
	"os"
	"testing"
	"time"

	"re-sep-scraper/scraper"
)

func createTempFile() (string, error) {
	file, err := os.CreateTemp("", "re-sep-test-*")
	if err != nil {
		return "", nil
	}
	file.Close()

	return file.Name(), nil
}

func TestCreateTable(t *testing.T) {
	tempfile, err := createTempFile()

	defer func() {
		defErr := os.Remove(tempfile)
		if defErr != nil {
			t.Error(err)
		}
	}()

	if err != nil {
		t.Fatal(err)
	}

	InitDB(tempfile)

	err = CreateTable()
	if err != nil {
		t.Fatalf("Table creation failed: %v", err)
	}

	db, err := connectDB()
	if err != nil {
		t.Fatalf("DB connection failed, %v", err)
	}
	defer db.Close()

	var result bool
	row := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM sqlite_schema WHERE type="table" AND name="articles")`)
	row.Scan(&result)

	if !result {
		t.Fatal("Table creation failed.")
	}
}

func TestInsertArticle(t *testing.T) {
	tempfile, err := createTempFile()

	defer func() {
		defErr := os.Remove(tempfile)
		if defErr != nil {
			t.Error(err)
		}
	}()

	if err != nil {
		t.Fatal(err)
	}

	InitDB(tempfile)

	err = CreateTable()
	if err != nil {
		t.Fatalf("Table creation failed: %v", err)
	}

	timeLayout := "2006-01-02"
	issueTime, _ := time.Parse(timeLayout, "2003-04-02")
	modifiedTime, _ := time.Parse(timeLayout, "2022-08-05")
	article := scraper.Article{
		EntryName: "foucault",
		Title:     "Michel Foucault",
		Author: []string{
			"Gutting, Gary",
			"Oksala, Johanna",
		},
		Issued:   issueTime,
		Modified: modifiedTime,
	}

	// Insert
	err = InsertArticle(article)
	if err != nil {
		t.Fatalf("Insert article failed: %v", err)
	}

	// Raw query check
	db, err := connectDB()
	if err != nil {
		t.Fatalf("DB connection failed, %v", err)
	}
	defer db.Close()

	var result string
	row := db.QueryRow(`SELECT id FROM articles`)
	err = row.Scan(&result)
	if err != nil {
		t.Fatal(err)
	}

	if result == "" {
		t.Fatal("Cannot find inserted article.")
	}
}
