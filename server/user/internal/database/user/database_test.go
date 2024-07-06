package database

import (
	"context"
	"testing"
)

func TestInsertUser(t *testing.T) {
	dbURL = "file:insertUserTest?mode=memory"
	InitUserDB()
	defer db.Close()

	user := InsertUser(context.Background(), "sub", "name")
	if user == nil || user.ID.String() == "" {
		t.Fatal("User insertion failed")
	}
}

func TestGetUserByUniqueID(t *testing.T) {
	dbURL = "file:getUserTest?mode=memory"
	InitUserDB()
	defer db.Close()

	InsertUser(context.Background(), "sub", "name")

	user := GetUserByUniqueID(context.Background(), "sub")
	if user == nil || user.ID.String() == "" {
		t.Fatal("Token insertion failed")
	}
}

func TestUpdateUsername(t *testing.T) {
	dbURL = "file:getUserTest?mode=memory"
	InitUserDB()
	defer db.Close()

	user := InsertUser(context.Background(), "sub", "name")

	user = UpdateUsername(context.Background(), user.Sub, "different_name")
	if user == nil {
		t.Fatal("Username update failed")
	}
	if user.Name != "different_name" {
		t.Fatalf("Mismatched username. Expected %s but got %s instead", "different_name", user.Name)
	}
}
