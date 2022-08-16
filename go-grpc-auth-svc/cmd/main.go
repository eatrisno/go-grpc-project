package main

import (
	"fmt"
	"log"
	"net"

	"github.com/eatrisno/go-grpc-auth-svc/pkg/pb"
	"github.com/eatrisno/go-grpc-auth-svc/pkg/services"
	"github.com/eatrisno/go-grpc-auth-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	h := utils.DBInit(config.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       config.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: config.JWTExpirationHours,
	}

	lis, err := net.Listen("tcp", config.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", config.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

	grpcServer.GracefulStop()
}
