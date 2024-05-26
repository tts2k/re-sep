package system

import (
	"os"
	"strings"
)

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
}

type EnvConfig struct {
	DBURL     string
	HTTPPort  string
	JWTSecret string
	Google    OAuthConfig
}

func isRunningTest() bool {
	for _, arg := range os.Args {
		if strings.HasSuffix(arg, ".test") {
			return true
		}
	}
	return false
}

func mustHaveEnv(key string) string {
	value := os.Getenv(key)
	if value == "" && isRunningTest() {
		return "test"
	}

	if value == "" {
		panic("Missing environment variable: " + key)
	}

	return value
}

var config EnvConfig = EnvConfig{
	DBURL:     mustHaveEnv("DB_URL"),
	HTTPPort:  mustHaveEnv("HTTP_PORT"),
	JWTSecret: mustHaveEnv("JWT_SECRET"),
	Google: OAuthConfig{
		ClientID:     mustHaveEnv("GOOGLE_CLIENT_ID"),
		ClientSecret: mustHaveEnv("GOOGLE_CLIENT_SECRET"),
	},
}

func Config() EnvConfig {
	return config
}
