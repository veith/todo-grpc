package middleware

import (
	"fmt"
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var headerAuthorize = "authorization"

func AuthFromMD(ctx context.Context, expectedScheme string) (string, error) {
	var val string

	cookie := metautils.ExtractIncoming(ctx).Get("cookie")

	if cookie != "" {
		val = cookie[14 : len(cookie)-1]
	} else {
		val = metautils.ExtractIncoming(ctx).Get(headerAuthorize)
	}

	if val == "" {

		return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+expectedScheme)

	}
	splits := strings.SplitN(val, " ", 2)
	if len(splits) < 2 {
		return "", status.Errorf(codes.Unauthenticated, "Bad authorization string")
	}
	if strings.ToLower(splits[0]) != strings.ToLower(expectedScheme) {
		return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+expectedScheme)
	}
	return splits[1], nil
}

var ExampleAuthFunc = func(ctx context.Context) (context.Context, error) {
	token, err := AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	fmt.Println(ctx)

	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	//grpc_ctxtags.Extract(ctx).Set("auth.sub", userClaimFromToken(tokenInfo))
	newCtx := context.WithValue(ctx, "tokenInfo", tokenInfo)
	return newCtx, nil
}

func parseToken(token string) (string, error) {
	token = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJBY2Nlc3NUb2tlbiI6ImxldmVsMSIsImlzcyI6InZlaXRoIiwic3ViIjoidGFzayJ9.fRKXYBnbWTTAPUxFV1E040QRxaIMa0Voh_kreJHuZre_u2cGROvi0aa_okhxVI6KkXah6FC7Uvc-YMBl-focHA"
	pubKey, err := ioutil.ReadFile("./keys/sample_key.pub")
	if err != nil {
		log.Fatal("Error reading private key")
		os.Exit(1)
	}
	rsaPublic, _ := crypto.ParseRSAPublicKeyFromPEM(pubKey)
	jwt, err := jws.ParseJWT([]byte(token))
	if err != nil {
		fmt.Println(err)

	}

	// Validate token
	if err = jwt.Validate(rsaPublic, crypto.SigningMethodHS512); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Token is valid")
	}
	return "ein Token", nil
}
func userClaimFromToken(tokenInfo string) (claim string) {
	return "rollen??"
}

func chekauth(ctx context.Context) error {

	headers, _ := metadata.FromIncomingContext(ctx)
	if headers["grpcgateway-user-agent"][0] == "err" {
		var err error
		return status.Errorf(codes.Unauthenticated, "Data Error: %s", err)
	}
	return nil
}
