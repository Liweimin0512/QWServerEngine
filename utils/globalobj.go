package utils

import (
	"QWServerEngine/qinterface"
	"encoding/json"
	"io/ioutil"
)

/*
	存储关于服务器框架的全局参数，供其他模块使用
*/

type GlobalObj struct {
	TcpServer qinterface.IServer
	Host      string
	TcpPort   int
	Name      string

	Version        string
	MaxConn        int    // 当前服务器主机允许的最大连接数
	MaxPackageSize uint32 // 当前框架数据包最大值
}

/*
 定义一个全局对外的obj
*/
var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	//data, err := ioutil.ReadFile("conf/serverConfig.json")
	data, err := ioutil.ReadFile("examples/conf/serverConfig.json")
	if err != nil {
		panic(err)
	}
	// 将data数据解析
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

/*
	提供一个初始化方法，初始化当前的 GlobalObject
*/
func init() {
	// 未加载配置文件时的默认值
	GlobalObject = &GlobalObj{
		Name:           "QW Server APP",
		Version:        "V0.0",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	// 尝试从配置文件加载一些用户自定义参数
	GlobalObject.Reload()
}
