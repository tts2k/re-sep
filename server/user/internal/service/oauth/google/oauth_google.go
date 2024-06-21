package service

import (
	"context"
	"log/slog"
	"net/http"
	"path"
	"time"

	tokenDB "re-sep-user/internal/database/token"
	userDB "re-sep-user/internal/database/user"

	common "re-sep-user/internal/service/oauth/common"
	config "re-sep-user/internal/system/config"
	authUtils "re-sep-user/internal/utils/auth"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type googleOIDC struct {
	provider    *oidc.Provider
	verifier    *oidc.IDTokenVerifier
	oAuthConfig *oauth2.Config
	name        string
}

var google googleOIDC

type OAuthStrategy interface {
	Login(w http.ResponseWriter, r *http.Request)
	Callback(w http.ResponseWriter, r *http.Request)
}

func init() {
	systemConfig := config.Config()
	google = googleOIDC{name: "google"}

	oidcProvider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		panic(err)
	}
	google.provider = oidcProvider

	google.oAuthConfig = &oauth2.Config{
		ClientID:     systemConfig.Google.ClientID,
		ClientSecret: systemConfig.Google.ClientSecret,
		RedirectURL:  path.Join(systemConfig.HTTPURL, "/oauth-callback/google"),
		Scopes:       []string{oidc.ScopeOpenID},
		Endpoint:     google.provider.Endpoint(),
	}
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	// Create state cookie
	oAuthState := authUtils.RandString(16)
	nonce := authUtils.RandString(16)

	authUtils.SetCallbackCookie(w, r, "state", oAuthState)
	authUtils.SetCallbackCookie(w, r, "nonce", nonce)

	oAuthURL := google.oAuthConfig.AuthCodeURL(oAuthState, oidc.Nonce(nonce))
	http.Redirect(w, r, oAuthURL, http.StatusTemporaryRedirect)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	oAuthState, err := r.Cookie("state")
	if err != nil {
		slog.Error("Cannot find state cookie", "error", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	state := r.FormValue("state")
	if state != oAuthState.Value {
		slog.Error("Mismatched state", "error", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	if code == "" {
		slog.Error("No auth code provided", "error", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Exchange token
	exchange, err := google.oAuthConfig.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		slog.Error("Could not exchange token", "error", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	rawIDToken, ok := exchange.Extra("id_token").(string)
	if !ok {
		slog.Error("No id_token field in oauth2 token", "error", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	idToken, err := google.verifier.Verify(context.Background(), rawIDToken)
	if err != nil {
		slog.Error("Cannot verify id token", "error", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	nonce, err := r.Cookie("nonce")
	if err != nil {
		slog.Error("Nonce not found", "error", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if idToken.Nonce != nonce.Value {
		slog.Error("Mismatched nonce", "error", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var claims struct {
		Sub string `json:"sub"`
	}
	idToken.Claims(&claims)

	user := userDB.GetUserByUniqueID(claims.Sub)
	if user == nil {
		slog.Warn("User not found. Creating new user", "error", err)
		user = userDB.InsertUser(google.name+":"+claims.Sub, common.DefaultUsername)
		if user == nil {
			slog.Error("User creation failed", "error", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
	}

	// Create 10 seconds token
	token := tokenDB.InsertToken(state, user.Sub, 10*time.Second)
	if token == nil {
		slog.Error("Token insertion failed", "error", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/?token="+state, http.StatusTemporaryRedirect)
}
