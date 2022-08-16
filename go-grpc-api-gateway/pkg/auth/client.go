package auth

import (
	"fmt"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/utils"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *utils.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
