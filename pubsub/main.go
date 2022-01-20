package main

import (
	proto "awesomeProject5/pubsub/proto"
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/metadata"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/asim/go-micro/v3/util/log"
)

type Sub struct{}

// Process Method can be of any name
func (s *Sub) Process(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

// Alternatively a function can be used
func subEv(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.2] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

func main() {

	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	// create a service
	service := micro.NewService(
		micro.Name("go.micro.srv.pubsub"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)
	// parse command line
	service.Init()

	// register subscriber
	if err := micro.RegisterSubscriber("example.topic,pubsub.1", service.Server(), new(Sub)); err != nil {
		log.Fatal(err)
	}

	// register subscriber with queue, each message is delivered to a unique subscriber
	if err := micro.RegisterSubscriber("example.topic.pubsub.2", service.Server(), subEv, server.SubscriberQueue("queue.pubsub")); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
