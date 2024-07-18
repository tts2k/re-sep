package store

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"re-sep-user/internal/database"
	token "re-sep-user/internal/database/token"
	user "re-sep-user/internal/database/user"
	pb "re-sep-user/internal/proto"
)

type AuthStore interface {
	RefreshToken(ctx context.Context, state string, duration time.Duration) (*pb.Token, error)
	GetTokenByState(ctx context.Context, state string) (*pb.Token, error)
	DeleteToken(ctx context.Context, state string) (*pb.Token, error)
	CleanTokens(ctx context.Context) (int64, error)
	InsertUser(ctx context.Context, sub string, neame string) (*pb.User, error)
	InsertToken(ctx context.Context, state string, userID string, duration time.Duration) (string, error)
	GetUserByUniqueID(ctx context.Context, id string) (*pb.User, error)
	UpdateUsername(ctx context.Context, sub string, newName string) (*pb.User, error)
	UpdateUserConfig(ctx context.Context, sub string, config *pb.UserConfig) (*pb.UserConfig, error)
	GetUserConfig(ctx context.Context, sub string) (*pb.UserConfig, error)
	Health() map[string]map[string]string
}

type authStore struct {
	userDB  *database.UserDB
	tokenDB *database.TokenDB
}

func NewAuthStore(userDB *database.UserDB, tokenDB *database.TokenDB) AuthStore {
	return &authStore{
		userDB:  userDB,
		tokenDB: tokenDB,
	}
}

func (s *authStore) InsertToken(ctx context.Context, state string, userID string, duration time.Duration) (string, error) {
	expires := time.Now().Add(duration)
	params := token.InsertTokenParams{
		State:   state,
		Userid:  userID,
		Expires: expires.Format(time.RFC3339),
	}

	result, err := s.tokenDB.Queries.InsertToken(ctx, params)
	if err != nil {
		slog.Error("InsertToken:", "error", err)
		return "", err
	}

	return result.State, nil
}

func (s *authStore) RefreshToken(ctx context.Context, state string, duration time.Duration) (*pb.Token, error) {
	expires := time.Now().Add(duration)
	params := token.RefreshTokenParams{
		State:   state,
		Expires: expires.Format(time.RFC3339),
	}

	result, err := s.tokenDB.Queries.RefreshToken(ctx, params)
	if err != nil {
		slog.Error("RefreshToken:", "error", err)
		return nil, err
	}

	expires, err = time.Parse(time.RFC3339, result.Expires)
	if err != nil {
		slog.Error("RefreshToken:", "Time parse error on expires", err)
		return nil, err
	}

	return &pb.Token{
		State:   result.State,
		UserId:  result.Userid,
		Expires: timestamppb.New(expires),
	}, nil
}

func (s *authStore) GetTokenByState(ctx context.Context, state string) (*pb.Token, error) {
	result, err := s.tokenDB.Queries.GetTokenByState(ctx, state)
	if err != nil {
		slog.Error("GetTokenByState:", "error", err)
		return nil, err
	}

	expires, err := time.Parse(time.RFC3339, result.Expires)
	if err != nil {
		slog.Error("GetTokenByState:", "error", err)
		return nil, err
	}

	return &pb.Token{
		State:   result.State,
		UserId:  result.Userid,
		Expires: timestamppb.New(expires),
	}, nil
}

func (s *authStore) DeleteToken(ctx context.Context, state string) (*pb.Token, error) {
	result, err := s.tokenDB.Queries.DeleteToken(ctx, state)
	if err != nil {
		slog.Error("DeleteToken:", "error", err)
		return nil, err
	}

	expires, err := time.Parse(time.RFC3339, result.Expires)
	if err != nil {
		slog.Error("GetTokenByState:", "error", err)
		return nil, err
	}

	return &pb.Token{
		State:   result.State,
		UserId:  result.Userid,
		Expires: timestamppb.New(expires),
	}, nil
}

func (s *authStore) CleanTokens(ctx context.Context) (int64, error) {
	count, err := s.tokenDB.Queries.CleanTokens(ctx)
	if err != nil {
		slog.Error("CleanTokens:", "error", err)
		return count, err
	}

	return count, nil
}

func (s *authStore) InsertUser(ctx context.Context, sub string, name string) (*pb.User, error) {
	params := user.InsertUserParams{
		ID:   uuid.New(),
		Name: name,
		Sub:  sub,
	}

	user, err := s.userDB.Queries.InsertUser(ctx, params)
	if err != nil {
		slog.Error("Cannot insert user", "database_error", err)
		return nil, err
	}

	return &pb.User{
		Sub:  user.Sub,
		Name: user.Name,
	}, nil
}

func (s *authStore) GetUserByUniqueID(ctx context.Context, id string) (*pb.User, error) {
	result, err := s.userDB.Queries.GetUserByUniqueID(ctx, id)
	if err != nil {
		slog.Error("Cannot get user by unique ID", "database_error", err)
		return nil, err
	}

	return &pb.User{
		Sub:  result.Sub,
		Name: result.Name,
	}, nil
}

func (s *authStore) UpdateUsername(ctx context.Context, sub string, newName string) (*pb.User, error) {
	params := user.UpdateUsernameParams{
		Name: newName,
		Sub:  sub,
	}

	result, err := s.userDB.Queries.UpdateUsername(ctx, params)

	return &pb.User{
		Sub:  result.Sub,
		Name: result.Name,
	}, err
}

func (s *authStore) UpdateUserConfig(ctx context.Context, sub string, config *pb.UserConfig) (*pb.UserConfig, error) {
	jsonConfig, err := json.Marshal(config)
	if err != nil {
		slog.Error("Cannot update username", "json_error", err)
		return nil, err
	}

	params := user.UpdateUserConfigParams{
		Sub:    sub,
		Config: string(jsonConfig),
	}

	uc, err := s.userDB.Queries.UpdateUserConfig(ctx, params)
	if err != nil {
		slog.Error("Cannot update user config", "database_error", err)
		return nil, err
	}

	var result pb.UserConfig
	err = json.Unmarshal([]byte(uc.Config), &result)
	if err != nil {
		slog.Error("Cannot parse user config", "json_error", err)
	}

	return &result, nil
}

func (s *authStore) GetUserConfig(ctx context.Context, sub string) (*pb.UserConfig, error) {
	result, err := s.userDB.Queries.GetUserConfig(ctx, sub)
	if err != nil {
		slog.Error("Cannot get user config", "database_error", err)
		return nil, err
	}

	var userConfig pb.UserConfig
	err = json.Unmarshal([]byte(result.Config), &userConfig)
	if err != nil {
		slog.Error("Cannot parse user config", "json_error", err)
	}

	return &userConfig, nil
}

func (s *authStore) Health() map[string]map[string]string {
	res := make(map[string]map[string]string)

	res["user"] = s.userDB.Health()
	res["token"] = s.tokenDB.Health()

	return res
}
