package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	config "re-sep-user/internal/system/config"
	"strings"
	"time"

	"github.com/go-jose/go-jose/v4"
	jwt "github.com/go-jose/go-jose/v4/jwt"
	"google.golang.org/grpc/metadata"
)

var systemConfig = config.Config()

func HTTPCall(method, rawURL, accessToken string) ([]byte, error) {
	httpClient := http.Client{
		Timeout: 30 * time.Second,
	}

	URL, _ := url.Parse(rawURL)
	queries := URL.Query()
	queries.Add("accessToken", accessToken)
	URL.RawQuery = queries.Encode()

	request := http.Request{
		Method: method,
		URL:    URL,
	}

	response, err := httpClient.Do(&request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func RandString(nByte int) string {
	b := make([]byte, nByte)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func SetCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}

func ExtractToken(ctx context.Context) (*jwt.Claims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}

	authHeader := md.Get("x-authorization")
	if len(authHeader) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	// Validate token
	authHeaderParts := strings.SplitN(authHeader[0], " ", 2)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return nil, fmt.Errorf("invalid authorization header")
	}

	// Decode the token
	token := authHeaderParts[1]
	tok, err := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})
	if err != nil {
		return nil, fmt.Errorf("cannot parse jwt token")
	}

	claims := jwt.Claims{}
	if err := tok.Claims([]byte(systemConfig.JWTSecret), &claims); err != nil {
		return nil, fmt.Errorf("cannot claim jwt token: %s", err)
	}

	return &claims, nil
}
