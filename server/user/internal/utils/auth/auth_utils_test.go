package utils

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	testUtils "re-sep-user/internal/utils/test"
	"testing"
)

func TestExtractToken(t *testing.T) {
	jwtToken, cl, err := testUtils.CreateJWTTestToken(systemConfig.JWTSecret)
	if err != nil {
		t.Fatal(err)
	}

	md := metadata.Pairs("x-authorization", fmt.Sprintf("Bearer %s", jwtToken))
	ctx := metadata.NewIncomingContext(context.Background(), md)

	claims, err := ExtractToken(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if claims.Subject != cl.Subject {
		t.Fatalf("Claim subject mismatch. Expected %s but got %s instead", cl.Subject, claims.Subject)
	}
}
