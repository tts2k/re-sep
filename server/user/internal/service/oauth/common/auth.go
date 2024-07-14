package service

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tokenDB "re-sep-user/internal/database/token"
	userDB "re-sep-user/internal/database/user"
	pb "re-sep-user/internal/proto"
	authUtils "re-sep-user/internal/utils/auth"
)

var waitGroup *sync.WaitGroup = &sync.WaitGroup{}

func PbAuth(ctx context.Context) (*pb.AuthResponse, error) {
	claims, err := authUtils.ExtractToken(ctx)
	if err != nil {
		slog.Error("Error extracting token", "authUtils.ExtractToken", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	// Get token
	token := tokenDB.GetTokenByState(context.Background(), claims.Subject)
	if token == nil {
		slog.Error("Error getting token", "tokenDB.GetTokenByState", nil)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	expires, err := time.Parse(time.RFC3339, token.Expires)
	if err != nil {
		slog.Error("Error parsing time", "time.Parse", nil)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	if time.Now().After(expires) {
		slog.Error("Token expired", "Token.Expires", expires)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	// Get user
	user := userDB.GetUserByUniqueID(ctx, token.Userid)
	if user == nil {
		slog.Error("Error getting user", "userDB.GetUserByUniqueID", nil)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	// Update token expiration
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		ctx := context.Background()

		token := tokenDB.RefreshToken(ctx, claims.Subject, 7*24*time.Hour)
		if token == nil {
			slog.Error("Error refreshing token", "tokenDB.RefreshToken", nil)
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

func PbGetUser(ctx context.Context) (*pb.User, error) {
	claims, err := authUtils.ExtractToken(ctx)
	if err != nil {
		slog.Error("Error extracting token", "authUtils.ExtractToken", err)
		return nil, err
	}

	// Get token
	token := tokenDB.GetTokenByState(ctx, claims.Subject)
	if token == nil {
		slog.Error("Error getting token", "tokenDB.GetUserByUniqueID", nil)
		return nil, err
	}

	// Get user
	user := userDB.GetUserByUniqueID(ctx, token.Userid)
	if user == nil {
		slog.Error("Error getting user", "userDB.GetUserByUniqueID", nil)
		return nil, err
	}

	return &pb.User{
		Name: user.Name,
		Sub:  user.Sub,
	}, nil
}
