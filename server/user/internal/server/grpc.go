package server

import (
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
