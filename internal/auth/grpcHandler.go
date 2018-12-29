package auth

import (
	"../proto"
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/oklog/ulid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
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

	// Todo: DB abfragen und Username mit Passwort gegen DB prüfen
	if req.Credentials.Password == "1234" {
		// erfolg
		_ = grpc.SendHeader(ctx, metadata.Pairs("Set-Cookie", "Authorization=Bearer "+string(createJWT())))
	} else {
		// falscher login
		return nil, status.Errorf(codes.Unauthenticated, "wrong password or username")

	}
	return &empty.Empty{}, nil
}

func (s *authServiceServer) Logout(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	_ = grpc.SendHeader(ctx, metadata.Pairs("Set-Cookie", "Authorization=deleted; expires=Thu, 01 Jan 1970 00:00:00 GMT"))
	return &empty.Empty{}, nil
}

func createJWT() string {
	bytes, err := ioutil.ReadFile("./keys/sample_key.priv")
	if err != nil {
		log.Fatal("Error reading private key")

	}
	rsaPrivate, keyErr := crypto.ParseRSAPrivateKeyFromPEM(bytes)
	if keyErr != nil {
		log.Fatal("Error parsing private key")
	}

	claims := jws.Claims{}
	claims.Set("AccessToken", "level1")
	claims.SetIssuer("veith")
	claims.SetSubject("task")
	now := time.Now()
	claims.SetIssuedAt(now)
	claims.SetExpiration(now.Add(time.Hour * time.Duration(24)))
	claims.SetNotBefore(now)
	claims.SetJWTID(GenerateUlidString())
	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)
	byteToken, err := jwt.Serialize(rsaPrivate)

	if err != nil {
		log.Fatal("Error signing the key. ", err)
		os.Exit(1)
	}

	return string(byteToken)
}

// Erzeuge eine ULID
func GenerateUlidString() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	newID, _ := ulid.New(ulid.Timestamp(t), entropy)
	return newID.String()
}
