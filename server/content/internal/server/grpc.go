package server

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"re-sep-content/internal/database"
	pb "re-sep-content/internal/proto"
)

type contentServer struct {
	pb.UnimplementedContentServer
	db database.Service
}

func newContentServer() *contentServer {
	return &contentServer{
		db: database.New(),
	}
}

func (s *contentServer) GetArticle(ctx context.Context, in *pb.EntryName) (*pb.Article, error) {
	if in.EntryName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: entryName")
	}

	slog.Info("Getting article: " + in.EntryName)
	return &pb.Article{}, nil
}
