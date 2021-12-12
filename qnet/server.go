package qnet

import (
	"fmt"
	"net"
	"projectS/qinterface"
)

// iServer接口实现，定义服务器模块
type Server struct {
	// 服务器名称
	ServerName string
	// 服务器绑定的IP版本
	IPVersion string
	// 服务器监听IP地址
	IP string
	// 服务器监听端口
	Port int
}

func (s *Server) Start() {
	fmt.Printf("[Start] server Listenner at IP :%s, PORT %d , is starting\n", s.IP, s.Port)

	go func() {
		// 获取TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addt error:", err)
			return
		}
		// 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "err ", err)
			return
		}

		fmt.Println("start QWServer success, ", s.ServerName, " succ, Listenning ... ")
		// 阻塞等待客户端链接， 处理客户端链接业务(读写）
		for {
			// 如果有客户端链接，阻塞返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 已经与客户端建立连接，处理业务逻辑
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}
					fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)
					// 回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()

		}
	}()
}

func (s *Server) Stop() {
	//TODO 将一些服务器资源、状态或者一些已经开辟的链接信息进行停止或回收
}

func (s *Server) Run() {

	// 启动server 的服务功能
	s.Start()

	// TODO 做一些启动服务器之后的业务逻辑

	// 阻塞状态
	select {}
}

/*
  初始化server模块的方法
*/
func Init(name string) qinterface.IServer {
	s := &Server{
		ServerName: name,
		IPVersion:  "tcp4",
		IP:         "0.0.0.0",
		Port:       8999,
	}
	return s
}
