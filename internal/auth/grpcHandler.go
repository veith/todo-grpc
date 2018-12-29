package auth

import (
	"../proto"
	"fmt"
	"github.com/SermoDigital/jose/jws"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"log"
	"os"
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

// Override Funktion um nicht über die default auth-middleware
func (s *authServiceServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (authServiceServer) Login(ctx context.Context, req *proto.CredentialsRequest) (*empty.Empty, error) {
	fmt.Println(ctx)
	token := string(createJWT())
	// create and send header
	header := metadata.Pairs("Set-Cookie", "Authorization=Bearer "+token)
	grpc.SendHeader(ctx, header)

	return &empty.Empty{}, nil
}

func createJWT() string {
	signKey, err := ioutil.ReadFile("./keys/sample_key.priv")
	if err != nil {
		log.Fatal("Error reading private key")
		os.Exit(1)
	}

	claims := jws.Claims{}
	claims.Set("AccessToken", "level1")
	claims.SetIssuer("veith")
	claims.SetSubject("task")
	signMethod := jws.GetSigningMethod("HS512")
	token := jws.NewJWT(claims, signMethod)
	byteToken, err := token.Serialize(signKey)
	if err != nil {
		log.Fatal("Error signing the key. ", err)
		os.Exit(1)
	}

	return string(byteToken)
}
