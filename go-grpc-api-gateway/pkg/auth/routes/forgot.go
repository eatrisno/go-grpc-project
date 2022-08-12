package routes

import (
	"context"
	"net/http"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type ForgotRequestBody struct {
	Email string `json:"email"`
}

func Forgot(ctx *gin.Context, c pb.AuthServiceClient) {
	b := ForgotRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Forgot(context.Background(), &pb.ForgotRequest{
		Email: b.Email,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
