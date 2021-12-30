package qinterface

/*
消息管理抽象层
 */

type IMsgHandle interface {
	// 调度/执行对应的Router消息处理方法
	DoMsgHandler(request IRequest)

	// 为消息添加具体处理逻辑
	AddRouter(msgID uint32, router IRouter)

	// 启动一个worker工作池
	StartWorkerPool()

	// 将消息交给 TaskQueue， 由 Worker 进行处理
	SendMsgToTaskQueue(request IRequest)
}