package service

import (
	"context"
	"log/slog"
	"net/http"
	"net/url"
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
var systemConfig config.EnvConfig = config.Config()

type OAuthStrategy interface {
	Login(w http.ResponseWriter, r *http.Request)
	Callback(w http.ResponseWriter, r *http.Request)
}

func init() {
	google = googleOIDC{name: "google"}

	oidcProvider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		panic(err)
	}
	google.provider = oidcProvider

	oidcConfig := &oidc.Config{
		ClientID: systemConfig.Google.ClientID,
	}
	google.verifier = oidcProvider.Verifier(oidcConfig)

	google.oAuthConfig = &oauth2.Config{
		ClientID:     systemConfig.Google.ClientID,
		ClientSecret: systemConfig.Google.ClientSecret,
		RedirectURL:  "http://" + path.Join(systemConfig.HTTPURL, "/oauth/google/callback"),
		Scopes:       []string{oidc.ScopeOpenID},
		Endpoint:     google.provider.Endpoint(),
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Create state cookie
	oAuthState, err := authUtils.RandString(16)
	if err != nil {
		slog.Error("Auth state generation failed", "authUtils.RandString", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}

	nonce, err := authUtils.RandString(16)
	if err != nil {
		slog.Error("Nonce generation failed", "authUtils.RandString", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}

	authUtils.SetCallbackCookie(w, r, "state", oAuthState)
	authUtils.SetCallbackCookie(w, r, "nonce", nonce)

	oAuthURL := google.oAuthConfig.AuthCodeURL(oAuthState, oidc.Nonce(nonce))
	http.Redirect(w, r, oAuthURL, http.StatusTemporaryRedirect)
}

func Callback(w http.ResponseWriter, r *http.Request) {
	oAuthState, err := r.Cookie("state")
	if err != nil {
		slog.Error("Cannot find state cookie", "error", err)
		http.Redirect(w, r, systemConfig.ClientURL, http.StatusTemporaryRedirect)
		return
	}

	state := r.FormValue("state")
	if state != oAuthState.Value {
		slog.Error("Mismatched state", "error", err)
		http.Redirect(w, r, systemConfig.ClientURL, http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	if code == "" {
		http.Redirect(w, r, systemConfig.ClientURL, http.StatusTemporaryRedirect)
		return
	}

	// Exchange token
	exchange, err := google.oAuthConfig.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		slog.Error("Could not exchange token", "error", err)
		http.Redirect(w, r, systemConfig.ClientURL, http.StatusTemporaryRedirect)
		return
	}

	rawIDToken, ok := exchange.Extra("id_token").(string)
	if !ok {
		slog.Error("No id_token field in oauth2 token", "error", err)
		http.Redirect(w, r, systemConfig.ClientURL, http.StatusTemporaryRedirect)
		return
	}
	idToken, err := google.verifier.Verify(context.Background(), rawIDToken)
	if err != nil {
		slog.Error("Cannot verify id token", "error", err)
		http.Redirect(w, r, systemConfig.ClientURL, http.StatusTemporaryRedirect)
		return
	}

	nonce, err := r.Cookie("nonce")
	if err != nil {
		slog.Error("Nonce not found", "error", err)
		http.Redirect(w, r, systemConfig.ClientURL, http.StatusTemporaryRedirect)
		return
	}
	if idToken.Nonce != nonce.Value {
		slog.Error("Mismatched nonce", "error", err)
		http.Redirect(w, r, systemConfig.ClientURL, http.StatusTemporaryRedirect)
		return
	}

	var claims struct {
		Sub string `json:"sub"`
	}
	err = idToken.Claims(&claims)
	if err != nil {
		slog.Warn("id token claim failed", "Claims", err)
	}

	user := userDB.GetUserByUniqueID(context.Background(), google.name+":"+claims.Sub)
	if user == nil {
		slog.Warn("User not found. Creating new user", "error", err)

		user = userDB.InsertUser(context.Background(), google.name+":"+claims.Sub, common.DefaultUsername)
		if user == nil {
			slog.Error("User creation failed", "error", err)
			http.Redirect(w, r, systemConfig.ClientURL, http.StatusTemporaryRedirect)
			return
		}
	}

	// Create 10 seconds token
	token := tokenDB.InsertToken(context.Background(), state, user.Sub, 10*time.Second)
	if token == nil {
		slog.Error("Token insertion failed", "error", err)
		http.Redirect(w, r, systemConfig.ClientURL, http.StatusTemporaryRedirect)
		return
	}

	redirectURL, _ := url.Parse(systemConfig.ClientURL)
	q := redirectURL.Query()
	q.Set("token", state)
	redirectURL.RawQuery = q.Encode()

	http.Redirect(w, r, redirectURL.String(), http.StatusTemporaryRedirect)
}
