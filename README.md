# QWEngine

## 介绍
使用 golang + JavaScript(TypeScript) 开发的游戏服务器框架

## 软件架构
软件架构说明

### 0.5 消息封装

定义一个消息的结构Message

#### 属性

- 消息ID
- 消息长度
- 消息内容

#### 方法

- Setter
- Getter

#### pack模块

定义一个解决TCP粘包问题的封包拆包模块

- 针对Message进行TLV格式的封装
- 针对Message进行TLV格式的拆包

#### 消息封装集成

- 将 Message 添加到Request属性中
- 修改链接读取数据的机制，将单纯的读取byte 改为 拆包形式读取
- 给链接提供一个发包机制： 将发送的消息进行打包，再发送

### 0.6 多路由模式

消息管理模式 支持多路由业务api调度管理的

#### 属性

- 消息ID 和 router 对应关系 map

#### 方法

- 根据msgID索引路由方法
- 添加路由方法到map中

#### 集成框架

- 将server模块中的router属性改为handler，修改初始化方法
- 将connection模块中的router属性改为handler，修改初始化方法
- connection的之前调度Router的业务 替换为 handler调度，修改StartReader方法

## 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

## 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

## 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

## 任务清单

- [ ] 使用 golang 完成核心代码TCP版本
- [ ] 使用 golang 完成核心代码UCP版本
- [ ] 使用 golang 完成核心代码Websocket版本
- [ ] 使用 typescript 作为后端脚本语言
- [ ] 通过配置一键生成Unreal客户端SDK
- [ ] 通过配置一键生成Unity客户端SDK
- [ ] 通过配置一键生成数据库ORM（对象关系映射）
- [ ] 基于Entity-Component加购的服务器，方便扩展

## 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)

