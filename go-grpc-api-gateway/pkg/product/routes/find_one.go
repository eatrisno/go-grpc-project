package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

// @Summary FindProduct
// @ID FindProduct
// @Produce json
// @Success 200 {object} pb.FindOneResponse
// @Router /product/:id [get]
// @Param Body body pb.FindOneRequest true "The body to create a thing"
// @Security ApiKeyAuth
func FindOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
