package qnet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

func TestDataPack(t *testing.T) {
	// 创建socket TCP
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err:", err)
		return
	}
	// 从客户端读取数据，拆包处理
	go func() {
		for true {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("Server accept error", err)
			}

			go func(conn net.Conn) {
				// 处理客户端的请求
				// ---------> 拆包的过程 <----------
				dp := NewDataPack()
				for true {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error")
						break
					}

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("Server unpacke err", io.ErrUnexpectedEOF)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						// msg 有数据 需要进行第二次读取
						msg := msgHead.(*Message)
						msg.SetMsgData(make([]byte, msg.GetMsgLen()))

						//根据datalen长度再次从io流读取
						_, err := io.ReadFull(conn, msg.data)
						if err != nil {
							fmt.Println("server unpack data err:", err)
							return
						}

						// 完整的一个消息已经读取完毕
						fmt.Println("--> Recv MsgID:", msg.GetMsgID(), ", datalen = ", msg.GetMsgLen(), ", data = ", msg.GetMsgData())
					}

				}

			}(conn)
		}
	}()

	// 模拟客户端连接
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}

	dp := NewDataPack()

	msg1 := &Message{
		msgID:   1,
		dataLen: 6,
		data:    []byte{'s', 'e', 'r', 'v', 'e', 'r'},
	}

	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 error", err)
		return
	}

	msg2 := &Message{
		msgID:   1,
		dataLen: 5,
		data:    []byte{'h', 'e', 'l', 'l', '0'},
	}

	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 error", err)
		return
	}

	sendData1 = append(sendData1, sendData2...)
	conn.Write(sendData1)

	// 客户端阻塞
	select {
	case <-time.After(time.Second):
		return
	}
}
