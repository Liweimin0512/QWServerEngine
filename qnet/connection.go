package qnet

import (
	"QWServerEngine/qinterface"
	"errors"
	"fmt"
	"io"
	"net"
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

	// 告知当前链接已经退出/停止的 channel
	ExitChan chan bool

	// 无缓冲管道，用于读写Goroutine之间的消息通信
	msgChan chan []byte

	// 消息的管理MsgID 和对应的处理业务API关系
	MsgHandler qinterface.IMsgHandle
}

// 读业务
func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running!]")

	defer fmt.Println("[Reader is exit !] c.ConnID = ", c.ConnID, ", remote addr is ", c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端数据到buf，最大512字节
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("recv buf err", err)
		//	continue
		//}

		// 创建一个拆包解包的对象
		dp := NewDataPack()

		// 得到Msg Head 二进制流的 8个字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
			return
		}

		// 拆包，得到msgID 和 msgDatalen 放在msg 消息中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error", err)
			return
		}

		// 根据dataLen 再次读取data， 放在 msg.data 中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}
		msg.SetMsgData(data)

		req := Request{
			conn: c,
			msg:  msg,
		}

		// 执行注册的路由方法
		// 根据绑定好的MsgID 找到对应处理api业务执行
		go c.MsgHandler.DoMsgHandler(&req)

	}
}

/*
	写业务Goroutine, 专门给客户发消息的模块
*/
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Gortine is running!]")
	defer fmt.Println(c.GetRemoteAddr().String(), " [conn Writer exit!]", c.GetRemoteAddr().String())

	// 阻塞等待channel 的消息，进行写给客户端
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error, ", err, "Conn Writer exit!")
			}
		case <-c.ExitChan:
			// 代表Reader已经退出
			return
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.ConnID)
	// 启动从当前链接的读数据
	go c.StartReader()
	// 启动当前链接写数据业务
	go c.StartWriter()
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ", c.ConnID)

	// 如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// 关闭socket链接
	c.Conn.Close()

	// 通知 Writer模块 关闭
	c.ExitChan <- true

	// 回收资源
	close(c.ExitChan)
	close(c.msgChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg 将我们要发送给客户端的数据先封包，再发送
func (c *Connection) SendMsg(msgID uint32, data []byte) error {

	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}
	// 将data 进行封包 msgDataLen|MsgID|Data
	dp := NewDataPack()

	binaryMsg, err := dp.Pack(NewMessagePackage(msgID, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgID)
		return errors.New("Pack error msg")
	}

	// 将数据发送给客户端
	//if _, err := c.Conn.Write(binaryMsg); err != nil {
	//	fmt.Println("write msg id", msgID, "error : ", err)
	//	return errors.New("conn Write error")
	//}
	c.msgChan <- binaryMsg

	return nil
}

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, handle qinterface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: handle,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
	}
	return c
}
