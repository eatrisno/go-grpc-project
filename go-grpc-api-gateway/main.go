package main

import (
	"log"

	_ "github.com/eatrisno/go-grpc-api-gateway/docs"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/auth"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/order"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/product"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Go GRPC Project
// @version 3.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:3000
// @BasePath /
// @query.collection.format multi
func main() {
	c, err := utils.LoadConfig(".")
	if err != nil {
		log.Println("cannot load config:", err)
	}

	r := gin.Default()
	v1 := r.Group("v1")
	authSvc := *auth.RegisterRoutes(v1, &c)
	product.RegisterRoutes(v1, &c, &authSvc)
	order.RegisterRoutes(v1, &c, &authSvc)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(c.Port)
}
