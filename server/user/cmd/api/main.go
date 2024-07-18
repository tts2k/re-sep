package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"

	"google.golang.org/grpc"

	"re-sep-user/internal/database"
	"re-sep-user/internal/server"
	"re-sep-user/internal/store"

	pb "re-sep-user/internal/proto"
	config "re-sep-user/internal/system/config"
	logger "re-sep-user/internal/system/logger"
	task "re-sep-user/internal/system/task"
)

func main() {
	// Get config
	systemConfig := config.Config()

	logger.InitLogger()

	//init DBs
	tokenDB := database.NewTokenDB()
	userDB := database.NewUserDB()

	// Migration
	userDB.Migrate()
	tokenDB.Migrate()

	// init stores
	authStore := store.NewAuthStore(userDB, tokenDB)

	// init grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", systemConfig.GRPCPort))
	if err != nil {
		slog.Error("Error listening on gRPC port", "net.Listen", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, server.NewAuthServer(authStore))

	// init http server
	server := server.NewServer(authStore)

	// run the gRPC server
	go func() {
		slog.Info("gRPC server listening on", "port", systemConfig.GRPCPort)
		err = s.Serve(lis)
		if err != nil {
			slog.Error("Error serving gRPC", "s.Serve", err)
		}
	}()

	// run the http server
	go func() {
		slog.Info("HTTP server listening on", "port", systemConfig.HTTPPort)
		err = server.ListenAndServe()
		if err != nil {
			slog.Error("Error serving HTTP", "server.ListenAndServe()", err)
		}
	}()

	cleanTokens := func(ctx context.Context) error {
		_, err := tokenDB.Queries.CleanTokens(ctx)
		if err != nil {
			return err
		}
		return nil
	}

	go task.StartTask(context.Background(), cleanTokens, 24*time.Hour, "Clean tokens")

	select {}
}
