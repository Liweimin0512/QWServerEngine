package qinterface

/*
	路由抽象接口
	路由里的数据都是IRequest
*/
type IRouter interface {
	// 在处理conn 业务之前的钩子方法 Hook
	PreHandle(request IRequest)

	// 在处理conn 业务 主方法 Hook
	Handle(request IRequest)

	// 在处理conn 业务之后的钩子方法 Hook
	PostHandle(request IRequest)
}
