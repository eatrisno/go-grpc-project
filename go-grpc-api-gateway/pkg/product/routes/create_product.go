package routes

import (
	"context"
	"net/http"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

// @Summary CreateProduct
// @ID CreateProduct
// @Produce json
// @Success 200 {object} pb.CreateProductResponse
// @Router /product [post]
// @Param Body body pb.CreateProductRequest true "The body to create a thing"
// @Security ApiKeyAuth
func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	var body pb.CreateProductRequest

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
