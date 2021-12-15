package qnet

import (
	"QWServerEngine/qinterface"
	"fmt"
	"strconv"
)

/*
消息处理模块的实现
 */
type MsgHandle struct {
	// 存放 MsgID 对应的处理方法
	Apis map[uint32] qinterface.IRouter
}

func (m *MsgHandle) DoMsgHandler(request qinterface.IRequest) {
	//1 从 request 中找到 msgID
	handler, ok := m.Apis[request.GetMsgID()]
	if !ok{
		fmt.Println("api msgID = " , request.GetMsgID(), "is NOT FOUND! Need Register!")
	}
	//2 根据MsgID 调度对应router业务即可
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (m *MsgHandle) AddRouter(msgID uint32, router qinterface.IRouter) {
	// 判断当前ID已经添加则不需要操作
	if _, ok := m.Apis[msgID]; ok {
		// ID 已经注册
		panic("repeat api , msgID = " + strconv.Itoa(int( msgID)))
	}
	m.Apis[msgID] = router
	fmt.Println("add api msgID = ", msgID, "success!")
}

func NewMsgHandle() *MsgHandle{
	return &MsgHandle{
		Apis: make(map[uint32] qinterface.IRouter),
	}
}