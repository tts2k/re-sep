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

type DBConfig struct {
	Path  string
	URL   string
	Token string
}

type EnvConfig struct {
	UserDB    DBConfig
	TokenDB   DBConfig
	HTTPPort  string
	GRPCPort  string
	BaseURL   string
	HTTPURL   string
	ClientURL string
	JWTSecret string
	Google    OAuthConfig
}

func (c EnvConfig) ConstructDBPath(dbName, fileName string) string {
	if testing.Testing() {
		return ""
	}

	var dbPath string
	if dbName == "user" {
		dbPath = path.Join(c.UserDB.Path, fileName)
	} else {
		dbPath = path.Join(c.UserDB.Path, fileName)
	}

	return dbPath
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

func mustHaveEnvTestDefault(key string, defaultValue string) string {
	if testing.Testing() {
		return defaultValue
	}

	value := os.Getenv(key)
	if value == "" {
		panic("Missing environment variable: " + key)
	}

	return value
}

// Put env into a struct for lsp autocompletion
var config EnvConfig = EnvConfig{
	UserDB: DBConfig{
		Path:  mustHaveEnv("USER_DB_PATH"),
		URL:   mustHaveEnv("USER_DB_URL"),
		Token: mustHaveEnv("USER_DB_TOKEN"),
	},
	TokenDB: DBConfig{
		Path:  mustHaveEnv("USER_DB_PATH"),
		URL:   mustHaveEnv("USER_DB_URL"),
		Token: mustHaveEnv("USER_DB_TOKEN"),
	},
	HTTPPort:  mustHaveEnv("HTTP_PORT"),
	GRPCPort:  mustHaveEnv("GRPC_PORT"),
	BaseURL:   mustHaveEnv("BASE_URL"),
	ClientURL: mustHaveEnv("CLIENT_URL"),
	JWTSecret: mustHaveEnvTestDefault("JWT_SECRET", "hTZ66AYJgxylQJmiTXstSdhTuq2D3DUw"),
	Google: OAuthConfig{
		ClientID:     mustHaveEnv("GOOGLE_CLIENT_ID"),
		ClientSecret: mustHaveEnv("GOOGLE_CLIENT_SECRET"),
	},
}

func init() {
	httpURL := os.Getenv("HTTP_URL")
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
