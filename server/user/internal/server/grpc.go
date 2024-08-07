package server

import (
	"log"
	"log/slog"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "re-sep-user/internal/proto"

	authService "re-sep-user/internal/service/oauth/common"

	"re-sep-user/internal/store"

	"github.com/bufbuild/protovalidate-go"
)

var DefaultUserConfig = pb.UserConfig{
	Font:     "serif",
	FontSize: 3,
	Justify:  false,
	Margin: &pb.Margin{
		Left:  3,
		Right: 3,
	},
}

type AuthServer struct {
	pb.UnimplementedAuthServer
	authStore store.AuthStore
	validator *protovalidate.Validator
}

func NewAuthServer(authStore store.AuthStore) *AuthServer {
	validator, err := protovalidate.New()
	if err != nil {
		log.Fatal(err)
	}

	return &AuthServer{
		authStore: authStore,
		validator: validator,
	}
}

func (as *AuthServer) Auth(ctx context.Context, _ *pb.Empty) (*pb.AuthResponse, error) {
	return authService.PbAuth(ctx, as.authStore)
}

func (as *AuthServer) UpdateUsername(ctx context.Context, username *pb.Username) (*pb.User, error) {
	user, err := authService.PbGetUser(ctx, as.authStore)
	if err != nil {
		slog.Error("Get user from context failed", "authService.PbGetUser", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	err = as.validator.Validate(user)
	if err != nil {
		slog.Error("Validation failed", "validator.Validate", err, "message", user)
		return nil, status.Error(codes.InvalidArgument, "Message validation failed")
	}

	dbUser, err := as.authStore.UpdateUsername(ctx, user.Sub, user.Name)
	if err != nil {
		slog.Error("Update username failed", "authStore.UpdateUsername", err)
		return nil, status.Error(codes.Internal, "Interal error")
	}

	return &pb.User{
		Name: dbUser.Name,
		Sub:  dbUser.Sub,
	}, nil
}

func (as *AuthServer) GetUserConfig(ctx context.Context, _ *pb.Empty) (*pb.UserConfig, error) {
	user, err := authService.PbGetUser(ctx, as.authStore)
	if err != nil || user == nil {
		slog.Error("Get user from context failed", "authService.PbGetUser", err)
		return &DefaultUserConfig, nil
	}

	userConfig, err := as.authStore.GetUserConfig(ctx, user.Sub)
	if err != nil {
		slog.Error("Get user config failed", "userDB.GetUserConfig", err)
		return &DefaultUserConfig, nil
	}

	return userConfig, nil
}

func (as *AuthServer) UpdateUserConfig(ctx context.Context, uc *pb.UserConfig) (*pb.UserConfig, error) {
	user, err := authService.PbGetUser(ctx, as.authStore)

	if err != nil || user == nil {
		slog.Error("Get user from context failed", "authService.PbGetUser", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	err = as.validator.Validate(uc)
	if err != nil {
		slog.Error("Validation failed", "validator.Validate", err, "message", uc)
		return nil, status.Error(codes.InvalidArgument, "Message validation failed")
	}

	result, err := as.authStore.UpdateUserConfig(ctx, user.Sub, uc)
	if result == nil {
		slog.Error("Update user config failed", "userDB.UpdateUserConfig", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return result, nil
}
