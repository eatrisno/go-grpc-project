package main

import (
	"log"
	"os"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/auth"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/order"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/product"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r)
	product.RegisterRoutes(r, &authSvc)
	order.RegisterRoutes(r, &authSvc)

	r.Run(os.Getenv("PORT"))
}
