package database

import (
	"testing"
	"time"
)

func TestInsertUser(t *testing.T) {
	dbURL = "file:insertUserTest?mode=memory"
	InitDB()
	defer db.Close()

	user := InsertUser("sub", "name")
	if user == nil || user.ID.String() == "" {
		t.Fatal("User insertion failed")
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

func TestInsertToken(t *testing.T) {
	dbURL = "file:insertTest?mode=memory"
	InitDB()
	defer db.Close()

	user := InsertUser("sub", "testUser")
	if user == nil {
		t.Fatal("Insert user failed")
	}

	t.Run("insert", func(t *testing.T) {
		token := InsertToken("test", "sub", 1*time.Second)
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
		token := InsertToken("test", "sub", 1*time.Second)
		if token != nil {
			t.Fatalf("Expected error on dup")
		}
	})
}

func TestGetTokenByState(t *testing.T) {
	dbURL = "file:getTokenTest?mode=memory"
	InitDB()
	defer db.Close()

	InsertUser("sub", "testUser")
	InsertToken("test", "sub", 1*time.Second)

	token := GetTokenByState("test")
	if token.Userid != "sub" {
		t.Fatalf("Invalid state. Expected %s but got %s instead.", "sub", token.Userid)
	}
}
