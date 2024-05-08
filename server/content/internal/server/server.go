package server

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"

	"re-sep-content/internal/database"

	pb "re-sep-content/internal/proto"
)

type Server struct {
	db   database.Service
	port int
}

const DefaultPort = "5000"

func Serve() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	grpcServer := grpc.NewServer()
	pb.RegisterContentServer(grpcServer, newContentServer())

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		slog.Error("failed to listen: %v", err)
		return err
	}
	slog.Info(fmt.Sprintf("Listening on port %s", port))

	grpcServer.Serve(lis)

	return nil
}
