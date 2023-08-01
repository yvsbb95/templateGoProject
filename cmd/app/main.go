package main

import (
	"github.com/yvsbb95/templateGoProject/pkg/config"
	"github.com/yvsbb95/templateGoProject/pkg/infrastructure/controller/grpc"
	"log"
)

func main() {
	// Read config
	configuration, grpcConfig := config.LoadConfig()
	log.Println("Config from local env", configuration)

	// Run GRPC Server
	grpc.InitServer(grpcConfig)
}
