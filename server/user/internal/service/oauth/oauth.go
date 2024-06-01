package service

import "re-sep-user/internal/database"

func NewOAuthProvider(provider string, dbService *database.Service) OAuthStrategy {
	switch provider {
	case "google":
		return newOAuthGoogle(dbService)
	default:
		return nil
	}
}
