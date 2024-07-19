package server

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"re-sep-user/internal/database"
	"re-sep-user/internal/store"
	config "re-sep-user/internal/system/config"
	utils "re-sep-user/internal/utils/test"

	pb "re-sep-user/internal/proto"
)

var systemConfig config.EnvConfig = config.Config()

func createAuthStore() store.AuthStore {
	userDB := database.NewUserDBMemory()
	tokenDB := database.NewTokenDBMemory()
	userDB.Migrate()
	tokenDB.Migrate()

	return store.NewAuthStore(userDB, tokenDB)
}

func TestGetUserConfig(t *testing.T) {
	type TestCase = struct {
		err    error
		setup  func(ctx context.Context, authStore store.AuthStore)
		expect *pb.UserConfig
		name   string
		user   string
	}

	testCases := []TestCase{
		{
			name:   "No user",
			user:   "",
			expect: &DefaultUserConfig,
			err:    status.Error(codes.Unauthenticated, "Unauthenticated"),
		},
		{
			name:   "With user but no config",
			user:   "test",
			expect: &DefaultUserConfig,
			err:    nil,
		},
		{
			name: "With user and config",
			user: "test",
			setup: func(ctx context.Context, authStore store.AuthStore) {
				uc := pb.UserConfig{
					FontSize: 1,
					Font:     "sans-serif",
					Justify:  true,
					Margin: &pb.Margin{
						Left:  1,
						Right: 2,
					},
				}

				_, _ = authStore.InsertUser(ctx, "userId", "user")
				_, _ = authStore.InsertToken(ctx, "token", "userId", 1*time.Hour)
				_, _ = authStore.UpdateUserConfig(ctx, "userId", &uc)
			},
			expect: &pb.UserConfig{
				FontSize: 1,
				Font:     "sans-serif",
				Justify:  true,
				Margin: &pb.Margin{
					Left:  1,
					Right: 2,
				},
			},
			err: nil,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			authStore := createAuthStore()

			var md metadata.MD
			if v.user != "" {
				jwtToken, _, err := utils.CreateJWTTestToken(systemConfig.JWTSecret, "token")
				if err != nil {
					t.Fatal(err)
				}

				md = metadata.Pairs("x-authorization", fmt.Sprintf("Bearer %s", jwtToken))
			}
			ctx := metadata.NewIncomingContext(context.Background(), md)

			// Setup test data
			if v.setup != nil {
				v.setup(ctx, authStore)
			}

			authServer := NewAuthServer(authStore)
			pbUc, err := authServer.GetUserConfig(ctx, nil)
			if v.err == nil && err != nil {
				t.Fatal(err)
			}
			if v.err != nil && err != nil {
				expect := v.err.Error()
				result := err.Error()

				if !errors.Is(v.err, err) {
					t.Fatalf("Mismatched error.\n Expected: %s\n Got: %s\n", expect, result)
				}
			}

			if !reflect.DeepEqual(pbUc.ProtoReflect().Interface(), (*v.expect).ProtoReflect().Interface()) {
				t.Logf("%v", pbUc.ProtoReflect().Interface())
				t.Logf("%v", (*v.expect).ProtoReflect().Interface())
				t.Fatal("Mismatched result config")
			}
		})
	}
}

func TestUpdateUserConfig(t *testing.T) {
	type TestCase = struct {
		err    error
		expect *pb.UserConfig
		input  *pb.UserConfig
		name   string
		user   string
	}

	testCases := []TestCase{
		{
			name: "Unauthenticated",
			user: "",
			err:  status.Error(codes.Unauthenticated, "Unauthenticated"),
		},
		{
			name: "Authenticated",
			user: "test",
			input: &pb.UserConfig{
				FontSize: 1,
				Justify:  true,
				Font:     "sans-serif",
				Margin: &pb.Margin{
					Left:  1,
					Right: 1,
				},
			},
			expect: &pb.UserConfig{
				FontSize: 1,
				Justify:  true,
				Font:     "sans-serif",
				Margin: &pb.Margin{
					Left:  1,
					Right: 1,
				},
			},
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			authStore := createAuthStore()

			var md metadata.MD
			if v.user != "" {
				jwtToken, _, err := utils.CreateJWTTestToken(systemConfig.JWTSecret, "token")
				if err != nil {
					t.Fatal(err)
				}

				md = metadata.Pairs("x-authorization", fmt.Sprintf("Bearer %s", jwtToken))
			}
			ctx := metadata.NewIncomingContext(context.Background(), md)

			authServer := NewAuthServer(authStore)
			pbUc, err := authServer.UpdateUserConfig(ctx, v.input)
			if v.err == nil && err != nil {
				t.Fatal(err)
			}
			if v.err != nil && err != nil {
				expect := v.err.Error()
				result := err.Error()

				if !errors.Is(v.err, err) {
					t.Fatalf("Mismatched error.\n Expected: %s\n Got: %s\n", expect, result)
				}
			}

			if !reflect.DeepEqual(pbUc, v.expect) {
				t.Fatal("Mismatched result config")
			}
		})
	}
}
