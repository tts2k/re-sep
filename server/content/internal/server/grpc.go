package server

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"re-sep-content/internal/database"
	pb "re-sep-content/internal/proto"
	"re-sep-content/internal/server/utils"
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

	// Get article
	article, err := s.db.GetArticle(ctx, in.EntryName)
	if article == nil || err != nil {
		slog.Error("Get article from context failed", "s.db.GetArticle", err)
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

	// Get User config
	ucCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	var userConfig *pb.UserConfig
	userConfig, err = s.authClient.GetUserConfig(ucCtx, nil)
	if userConfig == nil && err != nil {
		slog.Warn("Get user config failed", "s.authClient.GetUserConfig", err)
		userConfig = &DefaultUserConfig
	}
	if err != nil {
		slog.Warn("Get user config success with error", "s.authClient.GetUserConfig", err)
	}

	// Uncompress
	gzr, err := gzip.NewReader(bytes.NewBuffer(article.HTMLText))
	if err != nil {
		slog.Error("Uncompress content failed", "s.authClient.GetUserConfig", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	htmlText, err := io.ReadAll(gzr)
	if err != nil {
		slog.Error("Uncompress content failed", "io.ReadAll", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	resultHTML := utils.ApplyTemplate(string(htmlText), userConfig)

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
