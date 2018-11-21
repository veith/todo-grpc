package main

import (
	"../../protos"
	t "../../task"
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
	defer dbSession.Close() // Remember to close the database session.

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// DB session weitergeben
	t.ConnectDatabase(dbSession)

	// weitere Services kann man hier registrieren
	task.RegisterTodoServiceServer(grpcServer, t.GetServiceServer())

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
