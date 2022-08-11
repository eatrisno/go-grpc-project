package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/eatrisno/go-grpc-order-svc/pkg/client"
	"github.com/eatrisno/go-grpc-order-svc/pkg/db"
	"github.com/eatrisno/go-grpc-order-svc/pkg/pb"
	"github.com/eatrisno/go-grpc-order-svc/pkg/services"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	h := db.Init(os.Getenv("DB_URL"))

	lis, err := net.Listen("tcp", os.Getenv("PORT"))

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	productSvc := client.InitProductServiceClient(os.Getenv("PRODUCT_SVC_URL"))

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", os.Getenv("PORT"))

	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
