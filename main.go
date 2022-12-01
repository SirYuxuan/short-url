package main

import (
	"short-url/router"
	"short-url/support/redispool"
)

func main() {

	// 初始化Redis连接池
	redispool.Init()

	// 初始化服务路由
	server := router.InitRouter()

	if err := server.Run(":10003"); err != nil {
		panic("启动服务失败")
	}
}
