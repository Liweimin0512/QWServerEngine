package qnet

import "QWServerEngine/qinterface"

type Request struct {
	// 已经和客户端建立好的链接
	conn qinterface.IConnection
	// 客户端请求的数据
	data []byte
}

func (r Request) GetConnection() qinterface.IConnection {
	return r.conn
}

func (r Request) GetData() []byte {
	return r.data
}


