package qnet

type Message struct {
	msgID   uint32
	dataLen uint32
	data    []byte
}

func (m *Message) GetMsgID() uint32 {
	return m.msgID
}

func (m *Message) GetMsgLen() uint32 {
	return m.dataLen
}

func (m *Message) GetMsgData() []byte {
	return m.data
}

func (m *Message) SetMsgID(u uint32) {
	m.msgID = u
}

func (m *Message) SetMsgLen(u uint32) {
	m.dataLen = u
}

func (m *Message) SetMsgData(bytes []byte) {
	m.data = bytes
}
