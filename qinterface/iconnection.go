package qinterface

import "net"

/*
	定义链接模块的抽象层
*/
type IConnection interface {
	// 启动链接 让之前的链接准备开始工作
	Start()
	// 停止链接 结束当前链接的工作
	Stop()
	// 获取当前链接的绑定Socket conn
	GetTCPConnection() *net.TCPConn
	// 获取当前连接模块的链接ID
	GetConnID() uint32
	// 获取远程客户端的TCO状态IP Port
	GetRemoteAddr() net.Addr

	SendMsg(msgID uint32, data []byte) error
}

/*
	定义一个处理链接业务的方法
*/
type HandleFunc func(*net.TCPConn, []byte, int) error
