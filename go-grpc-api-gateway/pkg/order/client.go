package order

import (
	"fmt"
	"os"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient() pb.OrderServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(os.Getenv("ORDER_SVC_URL"), grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}
