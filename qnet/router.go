package qnet

import "QWServerEngine/qinterface"

/*
	实现router时，先嵌入此基类，然后根据需求对这个基类的方法进行重写
*/
type BaseRouter struct {
}

func (b *BaseRouter) PreHandle(request qinterface.IRequest) {
}

func (b *BaseRouter) Handle(request qinterface.IRequest) {
}

func (b *BaseRouter) PostHandle(request qinterface.IRequest) {
}
