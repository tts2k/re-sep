package database

import (
	"context"
	"encoding/json"
	"reflect"
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
	dbURL = "file:testUpdateUsername?mode=memory"
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

func TestUpdateUserConfig(t *testing.T) {
	dbURL = "file:testUpdateconfig?mode=memory"
	InitUserDB()
	defer db.Close()

	defConfString, _ := json.Marshal(defaultUserConfig)

	user := InsertUser(context.Background(), "sub", "name")
	config := UpdateUserConfig(context.Background(), user.Sub, defaultUserConfig)

	if config.Config != string(defConfString) {
		t.Fatal("Mismatched returned config string")
	}
}

func TestGetUserConfig(t *testing.T) {
	dbURL = "file:testUpdateconfig?mode=memory"
	InitUserDB()
	defer db.Close()

	user := InsertUser(context.Background(), "sub", "name")
	UpdateUserConfig(context.Background(), user.Sub, defaultUserConfig)

	config := GetUserConfig(context.Background(), user.Sub)
	if config == nil {
		t.Fatal("Get user config failed")
	}

	if !reflect.DeepEqual(*config, defaultUserConfig) {
		t.Fatal("Mismatched type")
	}
}
