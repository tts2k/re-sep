package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	tokenDB "re-sep-user/internal/database/token"
	userDB "re-sep-user/internal/database/user"
	googleOAuth "re-sep-user/internal/service/oauth/google"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/oauth/{provider}/login", s.handleOAuthLogin)
	mux.HandleFunc("/oauth/{provider}/callback", s.handleOAuthCallback)
	mux.HandleFunc("/oauth/logout", s.handleLogout)
	mux.HandleFunc("/health", s.healthHandler)

	return mux
}

func (s *Server) handleOAuthLogin(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("provider")
	if provider != "google" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	googleOAuth.Login(w, r)
}

func (s *Server) handleOAuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("provider")
	if provider != "google" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	googleOAuth.Callback(w, r)
}

func (s *Server) handleLogout(w http.ResponseWriter, r *http.Request) {
	state, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Not logged in", http.StatusBadRequest)
		return
	}

	c := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(w, c)

	token := tokenDB.DeleteToken(context.Background(), state.Value)
	if token.State == "" {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]map[string]string)

	res["user"] = userDB.Health()
	res["token"] = tokenDB.Health()

	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonRes)
}
