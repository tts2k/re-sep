package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	tokenDB "re-sep-user/internal/database/token"
	userDB "re-sep-user/internal/database/user"
	config "re-sep-user/internal/system/config"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
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
		t.Fatal("Cannot create user")
	}
}

func TestAuth(t *testing.T) {
	systemConfig := config.Config()
	initTestDB(t)
	key := []byte(systemConfig.JWTSecret)

	sig, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.HS256, Key: key},
		(&jose.SignerOptions{}).WithType("JWT"),
	)
	if err != nil {
		t.Fatal(err)
	}

	cl := jwt.Claims{
		Subject:  "token",
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Expiry:   jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	}
	raw, err := jwt.Signed(sig).Claims(cl).Serialize()
	if err != nil {
		t.Fatal(err)
	}

	md := metadata.Pairs("x-authorization", fmt.Sprintf("Bearer %s", raw))
	ctx := metadata.NewIncomingContext(context.Background(), md)

	res, err := PbAuth(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if res.Token != cl.Subject {
		t.Fatalf("Mismatched token. Expected %s but got %s instead.", cl.Subject, res.Token)
	}

	if res.User.Sub != "test" {
		t.Fatalf("Mismatched user. Expected user with sub %s but got %s instead.", "test", res.User.Sub)
	}

}
