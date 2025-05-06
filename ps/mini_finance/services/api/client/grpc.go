package client

import (
	"log"

	"github.com/gpjen/go_basic/ps/mini_finance/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var UserClient user.UserServiceClient

func InitGRPC() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	UserClient = user.NewUserServiceClient(conn)
}
