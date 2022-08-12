package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/eatrisno/go-grpc-auth-svc/pkg/db"
	"github.com/eatrisno/go-grpc-auth-svc/pkg/pb"
	"github.com/eatrisno/go-grpc-auth-svc/pkg/services"
	"github.com/eatrisno/go-grpc-auth-svc/pkg/utils"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	h := db.Init(os.Getenv("DB_URL"))

	jwt := utils.JwtWrapper{
		SecretKey:       os.Getenv("JWT_SECRET_KEY"),
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 30,
	}

	lis, err := net.Listen("tcp", os.Getenv("PORT"))

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", os.Getenv("PORT"))

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
