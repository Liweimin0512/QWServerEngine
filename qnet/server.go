package qnet

import (
	"QWServerEngine/qinterface"
	"QWServerEngine/utils"
	"errors"
	"fmt"
	"net"
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

	MsgHandler qinterface.IMsgHandle
}

/*
	TODO 定义当前客户端链接的所绑定的Handle API(暂时写死）
*/
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显业务
	fmt.Println("[Conn Handle] CallbackToClient ...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallbackToClient error")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("[server Info] ServerName : %s , listenner at ip : %s, port: %d is starting \n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[server Info] ServerVersion : %s , Max Connention : %d, Max Packeet Size: %d \n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	//fmt.Printf("[Start] server Listenner at IP :%s, PORT %d , is starting\n", s.IP, s.Port)

	go func() {
		// 开启消息队列及Worker工作池, 不需要go
		s.MsgHandler.StartWorkerPool()

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
		var cid uint32
		cid = 0

		// 阻塞等待客户端链接， 处理客户端链接业务(读写）
		for {
			// 如果有客户端链接，阻塞返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 将处理新链接的业务方法 和 conn 进行绑定 得到我们的链接模块
			dealConn := NewConnection(conn, cid, s.MsgHandler)
			cid++

			// 启动当前的链接业务处理
			go dealConn.Start()
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

func (s *Server) AddRouter(msgID uint32, router qinterface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Add Router Success!!!")
}

/*
  初始化server模块的方法
*/
func NewServer(name string) qinterface.IServer {
	s := &Server{
		ServerName: utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
	}
	return s
}
