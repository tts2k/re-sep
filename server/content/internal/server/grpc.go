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

	article := s.db.GetArticle(in.EntryName)
	if article == nil {
		return nil, fmt.Errorf("get article failed on entry name: %s", in.EntryName)
	}

	issuedTime, err := time.Parse(time.RFC3339, article.Issued)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("processing issued time failed on entry name: %s, ", in.EntryName)
	}

	modifiedTime, err := time.Parse(time.RFC3339, article.Modified)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("processing modified time failed on entry name: %s, ", in.EntryName)
	}

	protoArticle := pb.Article{
		Title:     article.Title,
		EntryName: article.EntryName,
		Issued:    timestamppb.New(issuedTime),
		Modified:  timestamppb.New(modifiedTime),
		HtmlText:  article.HTMLText,
	}

	err = json.Unmarshal([]byte(article.Author), &protoArticle.Authors)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("json unmarshalling author failed on entry name: %s, ", in.EntryName)
	}

	err = json.Unmarshal([]byte(article.TOC), &protoArticle.Toc)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("processing unmarshalling failed on entry name: %s, ", in.EntryName)
	}

	return &protoArticle, nil
}
