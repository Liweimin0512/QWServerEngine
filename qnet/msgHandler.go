package qnet

import (
	"QWServerEngine/qinterface"
	"QWServerEngine/utils"
	"fmt"
	"strconv"
)

/*
MsgHandle 消息处理模块的实现
 */
type MsgHandle struct {
	// 存放 MsgID 对应的处理方法
	Apis map[uint32] qinterface.IRouter

	// 负责 worker 读取任务的消息队列
	TaskQueue []chan qinterface.IRequest

	// 业务工作 worker 工作池的数量
	WorkerPoolSize uint32

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
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize, //从全局配置中获取
		TaskQueue: make([]chan qinterface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

// StartWorkerPool 启动一个worker工作池
func (m *MsgHandle)  StartWorkerPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		m.TaskQueue[i] = make(chan qinterface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		go m.startOneWorker(i, m.TaskQueue[i])
	}
}

// 启动一个工作流程
func (m *MsgHandle) startOneWorker(workerID int, taskQueue chan qinterface.IRequest)  {
	fmt.Println("Worker ID = ", workerID, "is Started...")

	for true {
		select {
		case request := <- taskQueue:
			m.DoMsgHandler(request)
		}
	}
}

// SendMsgToTaskQueue 将消息交给 TaskQueue， 由 Worker 进行处理
func (m *MsgHandle) SendMsgToTaskQueue(request qinterface.IRequest)  {
	// 1 平均分配
	workerID := request.GetConnection().GetConnID() % m.WorkerPoolSize
	fmt.Println(" Add ConnID = ", request.GetConnection().GetConnID(),
		"request MsgID = ", request.GetMsgID(), "to WorkerID = ", workerID)

	// 2 将消息发送给对应的worker的TaskQueue即可
	m.TaskQueue[workerID] <- request
}