package routes

import (
	"context"
	"net/http"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary login
// @ID login
// @Produce json
// @Success 200 {object} LoginRequestBody
// @Router /auth/login [post]
// @Param Body body LoginRequestBody true "The body to create a thing"
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
