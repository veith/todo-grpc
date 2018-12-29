package main

import (
	"../../internal/auth"
	"../../internal/proto"
	"../../internal/task"
	"./middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"upper.io/db.v3/sqlite"
)

const (
	port = ":9090"
)

// ConnectionURL implements a SQLite connection struct.
type ConnectionURL struct {
	Database string
	Options  map[string]string
}

var settings = sqlite.ConnectionURL{
	Database: `data/task.db`, // Path to database file.
}

func main() {

	dbSession, err := sqlite.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	defer dbSession.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_auth.UnaryServerInterceptor(middleware.ExampleAuthFunc),
	)))

	// DB session weitergeben
	task.ConnectDatabase(dbSession)

	// weitere Services kann man hier registrieren
	proto.RegisterAuthServiceServer(grpcServer, auth.GetServiceServer())
	proto.RegisterTaskServiceServer(grpcServer, task.GetServiceServer())

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
