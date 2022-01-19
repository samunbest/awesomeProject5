package handler

import (
	helloworld "awesomeProject5/helloworld/proto"
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
)

//demo
type demo struct{}

func NewDemo() *demo {
	return &demo{}
}

func (a *demo) InitRouter(router *gin.Engine) {
	router.POST("/demo", a.demo)
}

func (a *demo) demo(c *gin.Context) {
	// create a service

	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})

	service := micro.NewService(micro.Registry(consulRegistry), micro.Name("go.micro.srv.HelloWorld.client"))
	service.Init()

	client := helloworld.NewHelloworldService("go.micro.srv.HelloWorld", service.Client())

	rsp, err := client.Call(context.Background(), &helloworld.Request{
		Name: "world!",
	})
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": rsp.Msg})
}
