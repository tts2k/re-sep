package utils

import (
	"context"
	"fmt"
	"testing"
	"time"

	jose "github.com/go-jose/go-jose/v4"
	jwt "github.com/go-jose/go-jose/v4/jwt"
	"google.golang.org/grpc/metadata"
)

func TestExtractToken(t *testing.T) {
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

	claims, err := ExtractToken(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if claims.Subject != cl.Subject {
		t.Fatalf("Claim subject mismatch. Expected %s but got %s instead", cl.Subject, claims.Subject)
	}
}
