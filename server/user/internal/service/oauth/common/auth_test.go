package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	tokenDB "re-sep-user/internal/database/token"
	userDB "re-sep-user/internal/database/user"
	config "re-sep-user/internal/system/config"
	testUtils "re-sep-user/internal/utils/test"

	"google.golang.org/grpc/metadata"
)

func initTestDB(t *testing.T) {
	tokenDB.InitTokenDB()
	userDB.InitUserDB()

	user := userDB.InsertUser(context.Background(), "test", "tester")
	if user == nil {
		t.Fatal("Cannot create user")
	}
	token := tokenDB.InsertToken(context.Background(), "token", user.Sub, 10*time.Second)
	if token == nil {
		t.Fatal("Cannot create token")
	}
}

func TestAuth(t *testing.T) {
	systemConfig := config.Config()
	initTestDB(t)

	jwtToken, cl, err := testUtils.CreateJWTTestToken(systemConfig.JWTSecret)
	if err != nil {
		t.Fatal(err)
	}

	md := metadata.Pairs("x-authorization", fmt.Sprintf("Bearer %s", jwtToken))
	ctx := metadata.NewIncomingContext(context.Background(), md)

	res, err := PbAuth(ctx)
	if err != nil {
		t.Fatal(err)
	}

	waitGroup.Wait()

	if res.Token != cl.Subject {
		t.Fatalf("Mismatched token. Expected %s but got %s instead.", cl.Subject, res.Token)
	}

	if res.User.Sub != "test" {
		t.Fatalf("Mismatched user. Expected user with sub %s but got %s instead.", "test", res.User.Sub)
	}

	token := tokenDB.GetTokenByState(context.Background(), "token")
	if token == nil {
		t.Fatalf("Error getting token. Expected a token in database.")
	}

	expires, err := time.Parse(time.RFC3339, token.Expires)
	if err != nil {
		t.Fatal(err)
	}
	if !expires.After(time.Now().Add(10 * time.Second)) {
		t.Fatal("Token was not refreshed. Expected a token that last a week.")
	}
}
