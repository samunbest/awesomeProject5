package main

import (
	"awesomeProject5/registerConfiguration/config"
	"github.com/asim/go-micro/v3/logger"
)

func main() {
	// Register consul
	//reg := consul.NewRegistry(func(options *registry.Options) {
	//	options.Addrs = []string{"127.0.0.1:8500"}
	//})

	// 配置中心
	consulConfig, err := config.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		logger.Fatal(err)
	}

	// Mysql配置信息
	mysqlInfo, err := config.GetMysqlFromConsul(consulConfig, "mysql")
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Mysql配置信息:", mysqlInfo)
}
