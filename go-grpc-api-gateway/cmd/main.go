package main

import (
	"log"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/auth"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/order"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/product"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
