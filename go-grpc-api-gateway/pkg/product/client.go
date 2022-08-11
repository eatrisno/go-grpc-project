package product

import (
	"fmt"
	"os"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient() pb.ProductServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(os.Getenv("PRODUCT_SVC_URL"), grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}
