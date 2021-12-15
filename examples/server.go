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

// hello test 自定义路由
type HelloRouter struct {
	qnet.BaseRouter
}

// Test Handle
func (b *PingRouter) Handle(request qinterface.IRequest) {
	fmt.Println("call Router Handle")
	// 先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client: msgID = ", request.GetMsgID(),
		"data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

// Test Handle
func (b *HelloRouter) Handle(request qinterface.IRequest) {
	fmt.Println("call Hello Router Handle...")
	// 先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client: msgID = ", request.GetMsgID(),
		"data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("hello welcome to server game!!!"))
	if err != nil {
		fmt.Println(err)
	}
}

/*
	基于QWEngine开发的服务器端应用程序
*/
func main() {
	//1 创建一个 server 句柄，使用QW的API
	s := qnet.NewServer("MMO Game")
	//2 添加router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

	//3 启动服务器
	s.Run()
}
