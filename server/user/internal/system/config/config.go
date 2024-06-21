package system

import (
	"os"
	"path"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
}

type EnvConfig struct {
	DBPATH   string
	HTTPPort string
	BaseURL  string
	HTTPURL  string
	Google   OAuthConfig
}

func (c EnvConfig) ConstructDBPath(dbName string) string {
	return "file:" + path.Join(c.DBPATH, dbName)
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
	DBPATH:   mustHaveEnv("DB_PATH"),
	HTTPPort: mustHaveEnv("HTTP_PORT"),
	BaseURL:  mustHaveEnv("BASE_URL"),
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
		config.BaseURL = config.BaseURL[:len(config.BaseURL)-1]
	}
	httpURL = config.BaseURL + ":" + config.HTTPPort

	config.HTTPURL = httpURL
}

func Config() EnvConfig {
	return config
}
