package main

import (
	"fmt"
	"log"
	"net"

	"github.com/eatrisno/go-grpc-order-svc/pkg/client"
	"github.com/eatrisno/go-grpc-order-svc/pkg/db"
	"github.com/eatrisno/go-grpc-order-svc/pkg/pb"
	"github.com/eatrisno/go-grpc-order-svc/pkg/services"
	"github.com/eatrisno/go-grpc-order-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Println("cannot load config:", err)
	}

	h := db.Init(config.DBUrl)

	lis, err := net.Listen("tcp", config.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	productSvc := client.InitProductServiceClient(config.ProductServiceUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", config.Port)

	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

	grpcServer.GracefulStop()
}
