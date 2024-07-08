package server

import (
	"log/slog"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	userDB "re-sep-user/internal/database/user"
	pb "re-sep-user/internal/proto"
	authService "re-sep-user/internal/service/oauth/common"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
}

func (*AuthServer) Auth(ctx context.Context, _ *pb.Empty) (*pb.AuthResponse, error) {
	return authService.PbAuth(ctx)
}

func (*AuthServer) UpdateUsername(ctx context.Context, username *pb.Username) (*pb.User, error) {
	user, err := authService.PbGetUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	dbUser := userDB.UpdateUsername(ctx, user.Sub, user.Name)

	return &pb.User{
		Name: dbUser.Name,
		Sub:  dbUser.Sub,
	}, nil
}

func (*AuthServer) GetUserConfig(ctx context.Context, _ *pb.Empty) (*pb.UserConfig, error) {
	var result *pb.UserConfig

	getConfig := func(uc *userDB.UserConfig) {
		result = &pb.UserConfig{
			FontSize: uc.FontSize,
			Justify:  uc.Justify,
			Font:     uc.Font,
			Margin: &pb.Margin{
				Left:  uc.Margin.Left,
				Right: uc.Margin.Right,
			},
		}
	}

	user, err := authService.PbGetUser(ctx)
	if err != nil || user == nil {
		getConfig(&userDB.DefaultUserConfig)
		slog.Error("Get user from context failed", "authService.PbGetUser", err)
		println(result.Font)
		return result, nil
	}

	userConfig := userDB.GetUserConfig(ctx, user.Sub)
	if userConfig == nil {
		getConfig(&userDB.DefaultUserConfig)
		slog.Error("User has no config", "userDB.GetUserConfig", err)
		println(result.Font)
		return result, nil
	}

	getConfig(userConfig)
	return result, nil
}

func (*AuthServer) UpdateUserConfig(ctx context.Context, pbUC *pb.UserConfig) (*pb.UserConfig, error) {
	user, err := authService.PbGetUser(ctx)

	if err != nil || user == nil {
		slog.Error("Get user from context failed", "authService.PbGetUser", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	uc := userDB.UserConfig{
		FontSize: pbUC.FontSize,
		Justify:  pbUC.Justify,
		Font:     pbUC.Font,
		Margin: userDB.Margin{
			Left:  pbUC.Margin.Left,
			Right: pbUC.Margin.Right,
		},
	}
	result := userDB.UpdateUserConfig(ctx, user.Sub, &uc)
	if result == nil {
		slog.Error("Update user config failed", "userDB.UpdateUserConfig", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &pb.UserConfig{
			FontSize: result.FontSize,
			Justify:  result.Justify,
			Font:     result.Font,
			Margin: &pb.Margin{
				Left:  result.Margin.Left,
				Right: result.Margin.Right,
			},
		},
		nil
}
