package main

import (
	"QWServerEngine/qnet"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("client start ...")

	time.Sleep(1 * time.Second)

	// 直接链接远程服务器，得到一个connet
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err , exit!")
		return
	}
	for {
		// 发送封包消息 msgID = 0
		dp := qnet.NewDataPack()
		binaryMsg, err := dp.Pack(qnet.NewMessagePackage(0, []byte("test 0 v0.5 server client test message")))
		if err != nil {
			fmt.Println("Pack error", err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("Write error", err)
			return
		}
		// 服务器应该恢复msg数据，MsgID：1 pingpingping

		// 读取head 部分， 得到ID 和 dataLen
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error ", err)
			break
		}
		// 拆包
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack msgHead error ", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			// 根据DataLen进行第二次读取， 将data读出来
			msg := msgHead.(*qnet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error, ", err)
				return
			}

			fmt.Println("--> Recv Server Msg : ID = ", msg.ID, "Len = ", msg.DataLen, " , data = ", string(msg.Data))
		}

		// cpu 阻塞
		time.Sleep(1 * time.Second)
	}
}
