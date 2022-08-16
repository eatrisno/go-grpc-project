package product

import (
	"fmt"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/product/pb"
	"github.com/eatrisno/go-grpc-api-gateway/pkg/utils"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *utils.Config) pb.ProductServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}
