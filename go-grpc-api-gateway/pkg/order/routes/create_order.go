package routes

import (
	"context"
	"net/http"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/order/pb"
	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

// @Summary CreateOrder
// @ID CreateOrder
// @Produce json
// @Success 200 {object} CreateOrderRequestBody
// @Router /order [post]
// @Param Body body CreateOrderRequestBody true "The body to create a thing"
// @Security ApiKeyAuth
func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
