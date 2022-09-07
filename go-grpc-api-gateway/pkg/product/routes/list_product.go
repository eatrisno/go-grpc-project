package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

// @Summary ListProduct
// @ID ListProduct
// @Produce json
// @Success 200 {object} pb.ListProductResponse
// @Router /product/list [get]
// @Security ApiKeyAuth
func ListProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	res, err := c.ListProduct(context.Background(), &pb.ListProductRequest{
		Limit: int32(limit),
		Page:  int32(page),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
