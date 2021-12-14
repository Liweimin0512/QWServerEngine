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

// Test Handle
func (b *PingRouter) Handle(request qinterface.IRequest) {
	fmt.Println("call Router Handle")
	// 先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client: msgID = ", request.GetMsgID(),
		"data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(1,[]byte("ping...ping...ping"))
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
	s.AddRouter(&PingRouter{})
	//3 启动服务器
	s.Run()
}
