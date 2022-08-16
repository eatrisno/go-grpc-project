package order

import (
	"fmt"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/order/pb"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/utils"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *utils.Config) pb.OrderServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}
