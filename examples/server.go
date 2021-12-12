package main

import (
	"QWServerEngine/qinterface"
	"QWServerEngine/qnet"
	"fmt"
)

// ping test 自定义路由
type PingRouter struct {
	qnet.BaseRouter
}

// Test PreHandle
func (b *PingRouter) PreHandle(request qinterface.IRequest) {
	fmt.Println("call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

// Test Handle
func (b *PingRouter) Handle(request qinterface.IRequest) {
	fmt.Println("call Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ...ping ...ping ...\n"))
	if err != nil {
		fmt.Println("call back ping ...ping ...ping ... error")
	}
}

// Test PostHandle
func (b *PingRouter) PostHandle(request qinterface.IRequest) {
	fmt.Println("call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back after ping... error")
	}
}

/*
	基于QWEngine开发的服务器端应用程序
*/
func main() {
	//1 创建一个 server 句柄，使用QW的API
	s := qnet.Init("MMO Game")
	//2 添加router
	s.AddRouter(&PingRouter{})
	//3 启动服务器
	s.Run()
}
