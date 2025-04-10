package main

import (
	"context"
	"log"

	"github.com/gpjen/go_basic/ps/mini_finance/proto/user"
)

type UserHandler struct {
	user.UnimplementedUserServiceServer
}

func (h *UserHandler) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error) {
	log.Printf("Creating user: %s (%s)\n", req.Name, req.Email)

	// simulasi ID
	id := int32(1)

	return &user.UserResponse{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.UserResponse, error) {
	// simulasi setelah dapat data dari DB
	return &user.UserResponse{
		Id:    req.Id,
		Name:  "TEST 123",
		Email: "test@mail.com",
	}, nil
}
