package server

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "re-sep-content/internal/proto"
)

const DefaultPort = "5000"

func Serve() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	// GRPC client
	authURL := os.Getenv("AUTH_URL")
	if authURL == "" {
		log.Fatal("No auth URL")
	}
	conn, err :=
		grpc.NewClient(
			authURL,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
	if err != nil {
		log.Fatalf("Cannot connect to auth grpc server. Error: %v", authURL)
	}
	defer conn.Close()
	authClient := pb.NewAuthClient(conn)

	// GRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterContentServer(grpcServer, newContentServer(authClient))

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		slog.Error("failed to listen: %v", "net.Listen", err)
		return err
	}
	slog.Info(fmt.Sprintf("Listening on port %s", port))

	err = grpcServer.Serve(lis)
	if err != nil {
		slog.Error("Failed to start grpcServer", "grpcServer.Serve", err)
	}

	return nil
}
