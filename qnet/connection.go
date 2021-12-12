package qnet

import (
	"fmt"
	"net"
	"projectS/qinterface"
)

/*
	链接模块
*/
type Connection struct {
	// 当前链接socket TCP套接字
	Conn *net.TCPConn
	// 链接的ID
	ConnID uint32
	// 当前的链接状态
	isClosed bool
	// 当前链接所绑定的方法
	handleAPI qinterface.HandleFunc
	// 告知当前链接已经退出/停止的 channel
	ExitChan chan bool
}

// 读业务
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")

	defer fmt.Println("c.ConnID = ", c.ConnID, "Reader is exit, remote addr is ", c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端数据到buf，最大512字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		// 调用当前链接所绑定的HandleAPI
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID", c.ConnID, "Handle is error", err)
			break
		}
	}
}

// 写业务
func (c *Connection) StartWriter() {

}

func (c Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.ConnID)
	go c.StartReader()
	// 启动从当前链接的读数据
	// TODO 启动当前链接写数据业务

}

func (c Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ", c.ConnID)

	// 如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// 关闭socket链接
	c.Conn.Close()

	// 回收资源
	close(c.ExitChan)
}

func (c Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c Connection) Send(data []byte) error {
	return nil
}

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callback_api qinterface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callback_api,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}
