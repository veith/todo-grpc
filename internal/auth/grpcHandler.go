package auth

import (
	"../proto"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
)

var RegisterAuthServiceServer = proto.RegisterAuthServiceServer

// Gibt den grpc ServiceServer zurück
func GetServiceServer() proto.AuthServiceServer {
	var s authServiceServer
	return &s
}

// authServiceServer is used to implement authServiceServer.
type authServiceServer struct {
}

func (authServiceServer) Login(ctx context.Context, req *proto.CredentialsRequest) (*empty.Empty, error) {
	fmt.Println(ctx)

	return &empty.Empty{}, nil
}
