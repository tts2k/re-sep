package server

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	tokenDB "re-sep-user/internal/database/token"
	userDB "re-sep-user/internal/database/user"
	config "re-sep-user/internal/system/config"
	utils "re-sep-user/internal/utils/test"
)

var systemConfig config.EnvConfig = config.Config()

func initTestDB(t *testing.T) {
	tokenDB.InitTokenDB()
	userDB.InitUserDB()

	user := userDB.InsertUser(context.Background(), "test", "tester")
	if user == nil {
		t.Fatal("Cannot create user")
	}
	token := tokenDB.InsertToken(context.Background(), "token", user.Sub, 10*time.Second)
	if token == nil {
		t.Fatal("Cannot create token")
	}
}

func TestGetUserConfig(t *testing.T) {
	initTestDB(t)

	type TestCase = struct {
		err    error
		setup  func(ctx context.Context, t *testing.T)
		expect *userDB.UserConfig
		name   string
		user   string
	}

	testCases := []TestCase{
		{
			name:   "No user",
			user:   "",
			expect: &userDB.DefaultUserConfig,
			err:    status.Error(codes.Unauthenticated, "Unauthenticated"),
		},
		{
			name:   "With user but no config",
			user:   "test",
			expect: &userDB.DefaultUserConfig,
			err:    nil,
		},
		{
			name: "With user and config",
			user: "test",
			setup: func(ctx context.Context, t *testing.T) {
				uc := userDB.UserConfig{
					FontSize: 1,
					Font:     "sans-serif",
					Justify:  true,
					Margin: userDB.Margin{
						Left:  1,
						Right: 2,
					},
				}

				println("update")
				userDB.UpdateUserConfig(ctx, "test", uc)
			},
			expect: &userDB.UserConfig{
				FontSize: 1,
				Font:     "sans-serif",
				Justify:  true,
				Margin: userDB.Margin{
					Left:  1,
					Right: 2,
				},
			},
			err: nil,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			initTestDB(t)

			var md metadata.MD
			if v.user != "" {
				jwtToken, _, err := utils.CreateJWTTestToken(systemConfig.JWTSecret)
				if err != nil {
					t.Fatal(err)
				}

				md = metadata.Pairs("x-authorization", fmt.Sprintf("Bearer %s", jwtToken))
			}
			ctx := metadata.NewIncomingContext(context.Background(), md)

			// Setup test data
			if v.setup != nil {
				println(v.name)
				v.setup(ctx, t)
			}

			authServer := AuthServer{}
			pbUc, err := authServer.GetUserConfig(ctx, nil)
			if v.err == nil && err != nil {
				t.Fatal(err)
			}
			if v.err != nil && err != nil {
				expect := v.err.Error()
				result := err.Error()

				if expect != result {
					t.Fatalf("Mismatched error.\n Expected: %s\n Got: %s\n", expect, result)
				}
			}

			userConfig := userDB.UserConfig{
				FontSize: pbUc.FontSize,
				Justify:  pbUc.Justify,
				Font:     pbUc.Font,
				Margin: userDB.Margin{
					Left:  pbUc.Margin.Left,
					Right: pbUc.Margin.Right,
				},
			}

			t.Logf("%v", userConfig)

			if !reflect.DeepEqual(userConfig, *v.expect) {
				t.Fatal("Mismatched result config")
			}
		})
	}
}
