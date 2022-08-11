package auth

import (
	"fmt"
	"os"

	"github.com/eatrisno/go-grpc-api-gateway/pkg/auth/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient() pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(os.Getenv("AUTH_SVC_URL"), grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
