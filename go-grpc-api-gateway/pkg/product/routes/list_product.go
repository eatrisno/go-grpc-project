package routes

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func ListProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	switch {
	case limit > 50:
		limit = 50
	case limit < 10:
		limit = 10
	}
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page == 0 {
		page = 1
	}
	fmt.Println(limit, page)

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
