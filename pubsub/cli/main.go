package main

import (
	proto "awesomeProject5/pubsub/proto"
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/pborman/uuid"
	"time"
)

// send events using the publisher
func sendEv(topic string, p micro.Publisher) {
	t := time.NewTimer(time.Second)

	for _ = range t.C {
		// crate new event
		ev := &proto.Event{
			Id:        uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("Messaging you all day on %s", topic),
		}

		log.Logf("publishing %+v\n", ev)

		// publish an event
		if err := p.Publish(context.Background(), ev); err != nil {
			log.Logf("error publishing %v", err)
		}
	}
}

func main() {
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	serivce := micro.NewService(
		micro.Registry(consulRegistry),
		micro.Name("go.micro.cli.pubsub"),
	)
	serivce.Init()

	// create publisher
	pub1 := micro.NewEvent("example.topic.pubsub.1", serivce.Client())
	pub2 := micro.NewEvent("example.topic.pubsub.2", serivce.Client())

	// pub to topic 1
	go sendEv("example.topic.pubsub.1", pub1)
	// pub to topic 2
	go sendEv("example.topic.pubsub.2", pub2)

	// block forever
	select {}

}
