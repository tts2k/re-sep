package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"re-sep-content/internal/database"
	pb "re-sep-content/internal/proto"
	"re-sep-content/internal/server/utils"
)

type contentServer struct {
	pb.UnimplementedContentServer
	db         database.Service
	authClient pb.AuthClient
}

func newContentServer(authClient pb.AuthClient) *contentServer {
	return &contentServer{
		db:         database.New(),
		authClient: authClient,
	}
}

func (s *contentServer) GetArticle(ctx context.Context, in *pb.EntryName) (*pb.Article, error) {
	if in.EntryName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: entryName")
	}
	slog.Info("Getting article: " + in.EntryName)

	article := s.db.GetArticle(in.EntryName)
	if article == nil {
		slog.Error("Get article from context failed", "s.db.GetArticle", nil)
		return nil, status.Errorf(codes.NotFound, "get article failed on entry name: %s", in.EntryName)
	}

	issuedTime, err := time.Parse(time.RFC3339, article.Issued)
	if err != nil {
		slog.Error("Parse issued time failed", "time.Parse", err)
		return nil, status.Errorf(codes.Internal, "processing issued time failed on entry name: %s, ", in.EntryName)
	}

	modifiedTime, err := time.Parse(time.RFC3339, article.Modified)
	if err != nil {
		slog.Error("Parse issued time failed", "time.Parse", err)
		return nil, status.Errorf(codes.Internal, "processing modified time failed on entry name: %s, ", in.EntryName)
	}

	ucCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	userConfig, err := s.authClient.GetUserConfig(ucCtx, nil)
	if userConfig == nil && err != nil {
		slog.Error("Get user config failed", "s.authClient.GetUserConfig", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	if err != nil {
		slog.Warn("Get user config success with error", "s.authClient.GetUserConfig", err)
	}

	resultHTML := utils.ApplyTemplate(article.HTMLText, userConfig)

	protoArticle := pb.Article{
		Title:     article.Title,
		EntryName: article.EntryName,
		Issued:    timestamppb.New(issuedTime),
		Modified:  timestamppb.New(modifiedTime),
		HtmlText:  resultHTML,
	}

	err = json.Unmarshal([]byte(article.Author), &protoArticle.Authors)
	if err != nil {
		slog.Error("json unmarshal failed on article author", "json.Unmarshal", err)
		return nil, fmt.Errorf("json unmarshalling author failed on entry name: %s, ", in.EntryName)
	}

	err = json.Unmarshal([]byte(article.TOC), &protoArticle.Toc)
	if err != nil {
		slog.Error("json unmarshal failed on article TOC", "json.Unmarshal", err)
		return nil, fmt.Errorf("json unmarshalling failed on entry name: %s, ", in.EntryName)
	}

	return &protoArticle, nil
}
