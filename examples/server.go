package main

import "projectS/qnet"

/*
	基于QWEngine开发的服务器端应用程序
*/
func main() {
	// 创建一个 server 句柄，使用QW的API
	s := qnet.Init("MMO Game")
	s.Run()
}
