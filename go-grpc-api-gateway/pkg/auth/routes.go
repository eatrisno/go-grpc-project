package auth

import (
	"github.com/eatrisno/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *utils.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/forgot-password", svc.Forgot)
	routes.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) Forgot(ctx *gin.Context) {
	routes.Forgot(ctx, svc.Client)
}
