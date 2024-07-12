package database

import (
	"context"
	"reflect"
	"testing"
)

func TestInsertUser(t *testing.T) {
	InitUserDB()
	defer db.Close()

	user := InsertUser(context.Background(), "sub", "name")
	if user == nil || user.ID.String() == "" {
		t.Fatal("User insertion failed")
	}
}

func TestGetUserByUniqueID(t *testing.T) {
	InitUserDB()
	defer db.Close()

	InsertUser(context.Background(), "sub", "name")

	user := GetUserByUniqueID(context.Background(), "sub")
	if user == nil || user.ID.String() == "" {
		t.Fatal("Token insertion failed")
	}
}

func TestUpdateUsername(t *testing.T) {
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
	InitUserDB()
	defer db.Close()

	user := InsertUser(context.Background(), "sub", "name")
	config := UpdateUserConfig(context.Background(), user.Sub, &DefaultUserConfig)

	if !reflect.DeepEqual(*config, DefaultUserConfig) {
		t.Fatal("Mismatched returned config string")
	}
}

func TestGetUserConfig(t *testing.T) {
	InitUserDB()
	defer db.Close()

	user := InsertUser(context.Background(), "sub", "name")
	UpdateUserConfig(context.Background(), user.Sub, &DefaultUserConfig)

	config := GetUserConfig(context.Background(), user.Sub)
	if config == nil {
		t.Fatal("Get user config failed")
	}

	if !reflect.DeepEqual(*config, DefaultUserConfig) {
		t.Fatal("Mismatched type")
	}
}
