package qnet

import (
	"QWServerEngine/qinterface"
	"QWServerEngine/utils"
	"bytes"
	"encoding/binary"
	"errors"
)

type DataPack struct {
}

func (d DataPack) GetHeadLen() uint32 {
	// Datalen uint32 (4字节） + ID （4字节）
	return 8
}

func (d DataPack) Pack(msg qinterface.IMessage) ([]byte, error) {
	//创建一个存放字节流的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// 将 dataLen 、 MsgID 等写入 databuff 中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (d DataPack) Unpack(binaryData []byte) (qinterface.IMessage, error) {
	//创建一个存放字节流的缓冲
	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	//TODO 判断是否已经超出我们允许的最大长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too Large message data recv!")
	}

	return msg, nil
}

func NewDataPack() qinterface.IDataPack {
	return &DataPack{}
}
