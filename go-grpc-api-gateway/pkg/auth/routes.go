package auth

import (
	"github.com/eatrisno/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
