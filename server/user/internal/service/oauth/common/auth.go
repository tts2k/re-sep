package service

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "re-sep-user/internal/proto"
	"re-sep-user/internal/store"
	authUtils "re-sep-user/internal/utils/auth"
)

type OAuthStrategy interface {
	Login(w http.ResponseWriter, r *http.Request)
	Callback(w http.ResponseWriter, r *http.Request)
}

func PbAuth(ctx context.Context, authStore store.AuthStore) (*pb.AuthResponse, error) {
	claims, err := authUtils.ExtractToken(ctx)
	if err != nil {
		slog.Error("Error extracting token", "authUtils.ExtractToken", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	// Get token
	token, err := authStore.GetTokenByState(ctx, claims.Subject)
	if err != nil {
		slog.Error("Error getting token", "tokenDB.GetTokenByState", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	expires := token.Expires.AsTime()
	if time.Now().After(expires) {
		slog.Error("Token expired", "Token.Expires", expires)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	// Get user
	user, err := authStore.GetUserByUniqueID(ctx, token.UserId)
	if err != nil {
		slog.Error("Error getting user", "authStore.GetUserByUniqueID", nil)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	// Update token expiration
	go func() {
		ctx := context.Background()

		_, err := authStore.RefreshToken(ctx, claims.Subject, 7*24*time.Hour)
		if err != nil {
			slog.Error("Error refreshing token", "authStore.RefreshToken", err)
		}
	}()

	return &pb.AuthResponse{
		Token: claims.Subject,
		User: &pb.User{
			Sub:  user.Sub,
			Name: user.Name,
		},
	}, nil
}

func PbGetUser(ctx context.Context, authStore store.AuthStore) (*pb.User, error) {
	claims, err := authUtils.ExtractToken(ctx)
	if err != nil {
		slog.Error("Error extracting token", "authUtils.ExtractToken", err)
		return nil, err
	}

	// Get token
	token, err := authStore.GetTokenByState(ctx, claims.Subject)
	if token == nil {
		slog.Error("Error getting token", "tokenDB.GetUserByUniqueID", nil)
		return nil, err
	}

	// Get user
	user, err := authStore.GetUserByUniqueID(ctx, token.UserId)
	if err != nil {
		slog.Error("Error getting user", "userDB.GetUserByUniqueID", err)
		return nil, err
	}

	return &pb.User{
		Name: user.Name,
		Sub:  user.Sub,
	}, nil
}
