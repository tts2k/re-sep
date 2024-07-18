package server

import (
	"encoding/json"
	"log"
	"net/http"

	oauth "re-sep-user/internal/service/oauth/common"
	googleOAuth "re-sep-user/internal/service/oauth/google"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/oauth/{provider}/login", s.handleOAuthLogin)
	mux.HandleFunc("/oauth/{provider}/callback", s.handleOAuthCallback)
	mux.HandleFunc("/health", s.healthHandler)

	return mux
}

func (s *Server) handleOAuthLogin(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("provider")
	var oAuthStrategy oauth.OAuthStrategy

	if provider == "google" {
		oAuthStrategy = googleOAuth.NewGoogleOAuth(s.authStore)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	oAuthStrategy.Login(w, r)
}

func (s *Server) handleOAuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("provider")
	var oAuthStrategy oauth.OAuthStrategy

	if provider == "google" {
		oAuthStrategy = googleOAuth.NewGoogleOAuth(s.authStore)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	oAuthStrategy.Callback(w, r)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	health := s.authStore.Health()

	jsonRes, err := json.Marshal(health)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonRes)
}
