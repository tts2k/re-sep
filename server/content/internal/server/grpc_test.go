package server

import (
	"context"
	"testing"
	"time"

	"re-sep-content/internal/database"
	pb "re-sep-content/internal/proto"

	"google.golang.org/grpc"
)

type mockDatabaseService struct{}
type mockAuthClient struct {
	pb.AuthClient
}

func (mdb *mockDatabaseService) Health() map[string]string {
	return nil
}

func (mdb *mockDatabaseService) GetArticle(_ string) *database.Article {
	return &database.Article{
		EntryName: "test",
		Title:     "Test",
		Issued:    time.Now().Format(time.RFC3339),
		Modified:  time.Now().Format(time.RFC3339),
		HTMLText:  "<div>test<div>",
		Author:    `["Tester 1", "Tester 2"]`,
		TOC:       `[{"id":"Item 1","content":"Item 1"},{"id":"Item 2","content":"Item 2","subItems":[{"id":"Item 2-1","content":"Item 2-1"},{"id":"Item 2-2","content":"Item 2-2"}]},{"id":"Bib","content":"Bibliography","subItems":[{"id":"PrimSour","content":"Primary Sources"},{"id":"SecoSour","content":"Secondary Sources"}]},{"id":"Oth","content":"Other Internet Resources"},{"id":"Rel","content":"Related Entries"}]`,
	}
}

func (*mockAuthClient) Auth(ctx context.Context, in *pb.Empty, opts ...grpc.CallOption) (*pb.AuthResponse, error) {
	return nil, nil
}
func (*mockAuthClient) UpdateUsername(ctx context.Context, in *pb.Username, opts ...grpc.CallOption) (*pb.User, error) {
	return nil, nil
}
func (*mockAuthClient) UpdateUserConfig(ctx context.Context, in *pb.UserConfig, opts ...grpc.CallOption) (*pb.UserConfig, error) {
	return nil, nil
}
func (*mockAuthClient) GetUserConfig(ctx context.Context, in *pb.Empty, opts ...grpc.CallOption) (*pb.UserConfig, error) {
	return &pb.UserConfig{
		Font:     "serif",
		FontSize: 3,
		Justify:  false,
		Margin: &pb.Margin{
			Left:  3,
			Right: 3,
		},
	}, nil
}

func TestGetArticle(t *testing.T) {
	mockContentServer := contentServer{
		db:         &mockDatabaseService{},
		authClient: &mockAuthClient{},
	}

	article, err := mockContentServer.GetArticle(
		context.TODO(),
		&pb.EntryName{
			EntryName: "test",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if article == nil {
		t.Fatal("Error getting article. Nil returned.")
	}

	if len(article.Authors) != 2 {
		t.Fatalf("Mismatch author length. Expected %d but got %d instead.", 2, len(article.Authors))
	}

	if len(article.Toc) != 5 {
		t.Fatalf("Mismatch TOC length. Expected %d but got %d instead.", 5, len(article.Toc))
	}
}
