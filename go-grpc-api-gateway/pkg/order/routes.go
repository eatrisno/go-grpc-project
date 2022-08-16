package order

import (
	"github.com/eatrisno/go-grpc-api-gateway/pkg/auth"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/order/routes"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *utils.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
