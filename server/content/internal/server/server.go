package server

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"

	"re-sep-content/internal/database"

	pb "re-sep-content/internal/proto"
)

type Server struct {
	db   database.Service
	port int
}

const DefaultPort = 5000

func Serve() error {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = DefaultPort
	}

	grpcServer := grpc.NewServer()
	pb.RegisterContentServer(grpcServer, &contentServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		slog.Error("failed to listen: %v", err)
		return err
	}
	slog.Info(fmt.Sprintf("Listening on port %d", port))

	grpcServer.Serve(lis)

	return nil
}
