package service

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"re-sep-user/internal/store"
	config "re-sep-user/internal/system/config"
	testUtils "re-sep-user/internal/utils/test"

	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "re-sep-user/internal/proto"
)

func TestPbAuth(t *testing.T) {
	systemConfig := config.Config()
	mockCtrl := gomock.NewController(t)
	type TestCase = struct {
		ctx   context.Context
		err   error
		store func() store.AuthStore
		res   *pb.AuthResponse
		name  string
	}

	jwtToken, cl, err := testUtils.CreateJWTTestToken(systemConfig.JWTSecret)
	if err != nil {
		t.Fatal(err)
	}

	md := metadata.Pairs("x-authorization", fmt.Sprintf("Bearer %s", jwtToken))
	ctxWithAuth := metadata.NewIncomingContext(context.Background(), md)

	testCases := []TestCase{
		{
			name:  "No auth token",
			ctx:   context.Background(),
			store: nil,
			res:   nil,
			err:   status.Error(codes.Unauthenticated, "Unauthenticated"),
		},
		{
			name: "No token in db",
			ctx:  ctxWithAuth,
			store: func() store.AuthStore {
				m := store.NewMockAuthStore(mockCtrl)
				m.
					EXPECT().
					GetTokenByState(gomock.Eq(ctxWithAuth), gomock.Eq(cl.Subject)).
					Return(nil, errors.New(""))

				return m
			},
			res: nil,
			err: status.Error(codes.Unauthenticated, "Unauthenticated"),
		},
		{
			name: "Token expired",
			ctx:  ctxWithAuth,
			store: func() store.AuthStore {
				m := store.NewMockAuthStore(mockCtrl)
				m.
					EXPECT().
					GetTokenByState(gomock.Eq(ctxWithAuth), gomock.Eq(cl.Subject)).
					Return(&pb.Token{
						Expires: timestamppb.New(time.Now().AddDate(0, 0, -1)),
					}, nil)

				return m
			},
			res: nil,
			err: status.Error(codes.Unauthenticated, "Unauthenticated"),
		},
		{
			name: "Success",
			ctx:  ctxWithAuth,
			store: func() store.AuthStore {
				m := store.NewMockAuthStore(mockCtrl)
				m.
					EXPECT().
					GetTokenByState(gomock.Eq(ctxWithAuth), gomock.Eq(cl.Subject)).
					Return(&pb.Token{
						State:   cl.Subject,
						UserId:  "userId",
						Expires: timestamppb.New(time.Now().AddDate(0, 0, 1)),
					}, nil)

				m.
					EXPECT().
					GetUserByUniqueID(ctxWithAuth, gomock.Eq("userId")).
					Return(&pb.User{
						Sub:  "user",
						Name: "User",
					}, nil)

				return m
			},
			res: &pb.AuthResponse{
				Token: cl.Subject,
				User: &pb.User{
					Sub:  "user",
					Name: "User",
				},
			},
			err: nil,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			var authStore store.AuthStore
			if v.store != nil {
				authStore = v.store()
			}

			res, err := PbAuth(v.ctx, authStore)
			if !errors.Is(err, v.err) {
				t.Fatalf("Mismatched error.\nExpected: %v\nGot: %v\n", v.err, err)
			}

			if res == nil && v.res == nil {
				return
			} else if res == nil && v.res != nil {
				t.Fatalf("Mismatched res. Expected %s but got %v instead.", cl.Subject, nil)
			}

			if res.Token != v.res.Token {
				t.Fatalf("Mismatched token. Expected %s but got %s instead.", cl.Subject, res.Token)
			}

			if res.User.Sub != v.res.User.Sub {
				t.Fatalf("Mismatched user. Expected user with sub %s but got %s instead.", "test", res.User.Sub)
			}
		})
	}
}
