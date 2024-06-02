package service

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"path"
	"time"

	"re-sep-user/internal/database"
	"re-sep-user/internal/system"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthStrategy interface {
	Login(w http.ResponseWriter, r *http.Request)
	Callback(w http.ResponseWriter, r *http.Request)
}

type OAuthGoogle struct {
	config    oauth2.Config
	dbService *database.Service
}

func newOAuthGoogle(dbService *database.Service) *OAuthGoogle {
	systemConfig := system.Config()

	googleOAuthConfig := oauth2.Config{
		ClientID:     systemConfig.Google.ClientID,
		ClientSecret: systemConfig.Google.ClientSecret,
		RedirectURL:  path.Join(systemConfig.HTTPURL, "/oauth-callback/google"),
		Scopes:       []string{"openid"},
		Endpoint:     google.Endpoint,
	}

	return &OAuthGoogle{
		config:    googleOAuthConfig,
		dbService: dbService,
	}
}

func (o *OAuthGoogle) Login(w http.ResponseWriter, r *http.Request) {
	// Create state cookie
	oAuthState := generateStateOAuthCookie(w)

	oAuthURL := o.config.AuthCodeURL(oAuthState)
	http.Redirect(w, r, oAuthURL, http.StatusTemporaryRedirect)
}

func (o *OAuthGoogle) Callback(w http.ResponseWriter, r *http.Request) (string, error) {
	return "", nil
}

func (o *OAuthGoogle) GetUserUniqueID(accessToken string) (string, error) {
	return "", nil
}

func generateStateOAuthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	http.SetCookie(w, &http.Cookie{Name: "oauthState", Value: state, Expires: expiration})

	return state
}
