package qnet

import (
	"fmt"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err:", err)
		return
	}

	go func() {
		for true {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("Server accept error", err)
			}

			go func(conn net.Conn) {
				// 处理客户端的请求
				// ---------> 拆包的过程 <----------

			}(conn)
		}
	}()
}
