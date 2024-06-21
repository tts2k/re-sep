package database

import (
	"testing"
)

func TestInsertUser(t *testing.T) {
	dbURL = "file:insertUserTest?mode=memory"
	InitUserDB()
	defer db.Close()

	user := InsertUser("sub", "name")
	if user == nil || user.ID.String() == "" {
		t.Fatal("User insertion failed")
	}
}

func TestGetUserByUniqueID(t *testing.T) {
	dbURL = "file:getUserTest?mode=memory"
	InitUserDB()
	defer db.Close()

	InsertUser("sub", "name")

	user := GetUserByUniqueID("sub")
	if user == nil || user.ID.String() == "" {
		t.Fatal("Token insertion failed")
	}
}
