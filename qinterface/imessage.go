package qinterface

/*
请求消息封装，定义抽象接口
*/
type IMessage interface {
	GetMsgID() uint32
	GetMsgLen() uint32
	GetMsgData() []byte

	SetMsgID(uint32)
	SetMsgLen(uint32)
	SetMsgData([]byte)
}
