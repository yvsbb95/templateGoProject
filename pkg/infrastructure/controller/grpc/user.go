package grpc

import (
	"context"
	"github.com/yvsbb95/templateGoProject/pkg/infrastructure/proto"
	"log"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (r *UserServer) CreateUser(_ context.Context, request *pb.CreateUserRequest) (*pb.UserID, error) {
	log.Println("Create user GRPC call")
	log.Println(request.GetName(), request.GetEmail())
	return &pb.UserID{ID: 1}, nil
}

func (r *UserServer) GetUser(_ context.Context, request *pb.UserID) (*pb.User, error) {
	log.Println("Get user GRPC call")
	log.Println(request.GetID())
	return &pb.User{
		ID:    1,
		Name:  "Bhargav",
		Email: "yvsbb95@gmail.com",
	}, nil
}
