package grpc

import (
	"github.com/yvsbb95/templateGoProject/pkg/config"
	pb "github.com/yvsbb95/templateGoProject/pkg/infrastructure/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitServer(config *config.GrpcConfig) {
	// Run GRPC server
	lis, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
