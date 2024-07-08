package utils

import (
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

func CreateJWTTestToken(k string) (string, *jwt.Claims, error) {
	key := []byte(k)

	sig, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.HS256, Key: key},
		(&jose.SignerOptions{}).WithType("JWT"),
	)
	if err != nil {
		return "", nil, err
	}

	cl := jwt.Claims{
		Subject:  "token",
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Expiry:   jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	}
	raw, err := jwt.Signed(sig).Claims(cl).Serialize()
	if err != nil {
		return "", nil, err
	}

	return raw, &cl, nil
}
