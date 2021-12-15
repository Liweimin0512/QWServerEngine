package qnet

import "QWServerEngine/qinterface"

type Request struct {
	// 已经和客户端建立好的链接
	conn qinterface.IConnection
	// 客户端请求的数据
	msg qinterface.IMessage
}

func (r *Request) GetConnection() qinterface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMsgData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}