package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"

	"google.golang.org/grpc"

	tokenDB "re-sep-user/internal/database/token"
	userDB "re-sep-user/internal/database/user"
	"re-sep-user/internal/server"
	config "re-sep-user/internal/system/config"
	logger "re-sep-user/internal/system/logger"
	task "re-sep-user/internal/system/task"

	pb "re-sep-user/internal/proto"
)

func main() {
	// Get config
	systemConfig := config.Config()

	logger.InitLogger()

	//init DBs
	tokenDB.InitTokenDB()
	userDB.InitUserDB()

	// run the gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", systemConfig.GRPCPort))
	if err != nil {
		slog.Error("Error listening on gRPC port", "net.Listen", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &server.AuthServer{})
	server := server.NewServer()

	go func() {
		slog.Info("gRPC server listening on", "port", systemConfig.GRPCPort)
		err = s.Serve(lis)
		if err != nil {
			slog.Error("Error serving gRPC", "s.Serve", err)
		}
	}()

	go func() {
		slog.Info("HTTP server listening on", "port", systemConfig.HTTPPort)
		err = server.ListenAndServe()
		if err != nil {
			slog.Error("Error serving HTTP", "server.ListenAndServe()", err)
		}
	}()

	cleanTokens := func(ctx context.Context) error {
		tokenDB.CleanTokens(ctx)
		return nil
	}

	go task.StartTask(context.Background(), cleanTokens, 24*time.Hour, "Clean tokens")

	select {}
}
