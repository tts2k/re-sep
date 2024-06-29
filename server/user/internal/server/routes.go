package server

import (
	"encoding/json"
	"log"
	"net/http"

	tokenDB "re-sep-user/internal/database/token"
	userDB "re-sep-user/internal/database/user"
	googleOAuth "re-sep-user/internal/service/oauth/google"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/oauth/{provider}/login", s.handleOAuthLogin)
	mux.HandleFunc("/oauth/{provider}/callback", s.handleOAuthCallback)
	mux.HandleFunc("/health", s.healthHandler)

	return mux
}

func (s *Server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	print(r.PathValue("provider"))

	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
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

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]map[string]string)

	res["user"] = userDB.Health()
	res["token"] = tokenDB.Health()

	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonRes)
}
