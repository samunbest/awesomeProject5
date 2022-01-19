package main

import (
	"awesomeProject5/helloworld/handler"
	pb "awesomeProject5/helloworld/proto"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
)

const (
	ServerName = "go.micro.srv.HelloWorld" // server name
)

func main() {

	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	// Create service
	service := micro.NewService(
		micro.Name(ServerName),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)

	// Register handler
	if err := pb.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
