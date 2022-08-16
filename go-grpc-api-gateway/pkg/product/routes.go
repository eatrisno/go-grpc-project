package product

import (
	"github.com/eatrisno/go-grpc-api-gateway/pkg/auth"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/product/routes"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *utils.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
	routes.GET("/list", svc.ListProduct)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FineOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) ListProduct(ctx *gin.Context) {
	routes.ListProduct(ctx, svc.Client)
}
