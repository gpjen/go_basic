package main

import (
	"log"
	"net"

	"github.com/gpjen/go_basic/ps/mini_finance/proto/user"
	"google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	server := grpc.NewServer()
	user.RegisterUserServiceServer(server, &UserHandler{})

	log.Println("User service running at :50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
