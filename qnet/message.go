package qnet

type Message struct {
	ID      uint32
	DataLen uint32
	Data    []byte
}

func (m *Message) GetMsgID() uint32 {
	return m.ID
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) GetMsgData() []byte {
	return m.Data
}

func (m *Message) SetMsgID(u uint32) {
	m.ID = u
}

func (m *Message) SetMsgLen(u uint32) {
	m.DataLen = u
}

func (m *Message) SetMsgData(bytes []byte) {
	m.Data = bytes
}

func NewMessagePackage(id uint32, data []byte) *Message	{
	return &Message{
		ID: id,
		DataLen: uint32(len(data)),
		Data: data,
	}
}