package main

import (
	"fmt"
	"log"
	"net"

	"github.com/eatrisno/go-grpc-product-svc/pkg/db"
	"github.com/eatrisno/go-grpc-product-svc/pkg/pb"
	"github.com/eatrisno/go-grpc-product-svc/pkg/services"
	"github.com/eatrisno/go-grpc-product-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	h := db.Init(config.DBUrl)

	lis, err := net.Listen("tcp", config.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", config.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

	grpcServer.GracefulStop()
}
