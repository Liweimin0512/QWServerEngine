package qinterface

type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
	GetMsgID() uint32
}
