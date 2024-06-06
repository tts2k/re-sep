package database

import (
	"testing"
	"time"
)

func TestInsertToken(t *testing.T) {
	dbURL = "file:insertTest?mode=memory"
	InitDB()
	defer db.Close()

	t.Run("insert", func(t *testing.T) {
		token := InsertToken("test", "testToken", "refresh Token", 1)
		if token == nil {
			t.Fatal("Token insertion failed")
		}
		if token.State != "test" {
			t.Fatalf("Invalid state. Expected %s but got %s instead.", "test", token.State)
		}
		if token.Token != "testToken" {
			t.Fatalf("Invalid state. Expected %s but got %s instead.", "test", token.Token)
		}
	})

	t.Run("insert with duplicate", func(t *testing.T) {
		token := InsertToken("test", "testToken", "refresh Token", 1)
		if token != nil {
			t.Fatalf("Expected error on dup")
		}
	})
}

func TestGetTokenByState(t *testing.T) {
	dbURL = "file:getTokenTest?mode=memory"
	InitDB()
	defer db.Close()

	InsertToken("test", "testToken", "refresh token", 1*time.Second)

	token := GetTokenByState("test")
	if token != "testToken" {
		t.Fatalf("Invalid state. Expected %s but got %s instead.", "test", token)
	}
}

func TestInsertUser(t *testing.T) {
	dbURL = "file:insertUserTest?mode=memory"
	InitDB()
	defer db.Close()

	user := InsertUser("sub", "name")
	if user == nil || user.ID.String() == "" {
		t.Fatal("Token insertion failed")
	}
}

func TestGetUserByUniqueID(t *testing.T) {
	dbURL = "file:getUserTest?mode=memory"
	InitDB()
	defer db.Close()

	InsertUser("sub", "name")

	user := GetUserByUniqueID("sub")
	if user == nil || user.ID.String() == "" {
		t.Fatal("Token insertion failed")
	}
}
