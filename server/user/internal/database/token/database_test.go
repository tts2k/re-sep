package database

import (
	"context"
	"testing"
	"time"
)

func TestInsertToken(t *testing.T) {
	dbURL = "file:insertTest?mode=memory"
	InitTokenDB()
	defer db.Close()

	t.Run("insert", func(t *testing.T) {
		token := InsertToken(context.Background(), "test", "sub", 1*time.Second)
		if token == nil {
			t.Fatal("Token insertion failed")
		}
		if token.State != "test" {
			t.Fatalf("Invalid state. Expected %s but got %s instead.", "test", token.State)
		}
		if token.Userid != "sub" {
			t.Fatalf("Invalid state. Expected %s but got %s instead.", "sub", token.Userid)
		}
	})

	t.Run("insert with duplicate", func(t *testing.T) {
		token := InsertToken(context.Background(), "test", "sub", 1*time.Second)
		if token != nil {
			t.Fatalf("Expected error on dup")
		}
	})
}

func TestGetTokenByState(t *testing.T) {
	dbURL = "file:getTokenTest?mode=memory"
	InitTokenDB()
	defer db.Close()

	InsertToken(context.Background(), "test", "sub", 1*time.Second)

	token := GetTokenByState(context.Background(), "test")
	if token.Userid != "sub" {
		t.Fatalf("Invalid state. Expected %s but got %s instead.", "sub", token.Userid)
	}
}
