package main

import (
	helloworld "awesomeProject5/helloworld/proto"
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := helloworld.NewHelloworldService("go.micro.srv.HelloWorld", service.Client())

	// Make request
	rsp, err := cl.Call(context.Background(), &helloworld.Request{
		Name: "World!",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Msg)
}
