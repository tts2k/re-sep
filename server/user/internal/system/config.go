package system

import (
	"os"
	"testing"
)

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
}

type EnvConfig struct {
	DBURL     string
	HTTPPort  string
	BaseURL   string
	HTTPURL   string
	JWTSecret string
	Google    OAuthConfig
}

func mustHaveEnv(key string) string {
	if testing.Testing() {
		return "test"
	}

	value := os.Getenv(key)
	if value == "" {
		panic("Missing environment variable: " + key)
	}

	return value
}

// Put env into a struct for lsp autocompletion
var config EnvConfig = EnvConfig{
	DBURL:     mustHaveEnv("DB_URL"),
	HTTPPort:  mustHaveEnv("HTTP_PORT"),
	BaseURL:   mustHaveEnv("BASE_URL"),
	JWTSecret: mustHaveEnv("JWT_SECRET"),
	Google: OAuthConfig{
		ClientID:     mustHaveEnv("GOOGLE_CLIENT_ID"),
		ClientSecret: mustHaveEnv("GOOGLE_CLIENT_SECRET"),
	},
}

func init() {
	httpURL := os.Getenv("httpURL")
	if httpURL != "" {
		config.HTTPURL = httpURL
		return
	}

	// Get http url from base url and port if the value is not provided
	if config.BaseURL[len(config.BaseURL)-1] == '/' {
		httpURL = config.BaseURL[:len(config.BaseURL)-1] + ":" + config.HTTPPort
	} else {
		httpURL = config.BaseURL + ":" + config.HTTPPort
	}

	config.HTTPURL = httpURL
}

func Config() EnvConfig {
	return config
}
