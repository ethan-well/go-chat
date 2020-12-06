## 项目简介
这是使用 Golang 网络编程实现的一个多人在线聊天系统，使用 goroutine 达到高并发的效果，使用 redis 来保存用户的注册信息。
项目目录结构如下：
```
.
├── README.md
├── client // 客户端代码
│   ├── logger // 自定义的日志打印器
│   │   └── logger.go
│   ├── main.go // 主函数
│   ├── model // model 层
│   │   └── user.go
│   ├── process // 处理与服务器端的连接，收发消息
│   │   ├── messageProcess.go
│   │   ├── serverProcess.go
│   │   └── userProcess.go
│   └── utils
│       └── utils.go
├── common // 客户端和服务端公用代码，主要用来定义客户端和服务端之间通信约定的消息
│   └── message
│       └── message.go
├── config // 配置信息
│   ├── config.go
│   └── config.json
└── server // 服务端代码
    ├── main // 主函数
    │   ├── main.go
    │   └── redis.go
    ├── model // model 层
    │   ├── clientConn.go
    │   ├── error.go
    │   ├── user.go
    │   └── userDao.go
    ├── process // 处理与客户端的连接，收发消息
    │   ├── groupMessageProcess.go // 处理群消息
    │   ├── onlineInfoProcess.go // 显示在线用户
    │   ├── pointToPointMessageProcess.go // 处理点对点聊天消息
    │   ├── processor.go // 消息处理器入口
    │   └── userProcess.go // 处理和用户登陆注册相关消息
    └── utils
        └── utils.go
```

服务端和客户端代码基本独立，server 目录下是服务端代码，client 目录下是客户端代码，common 目录下的包由服务端和客户端共同使用

## 本地运行本项目(Unix 系统下)
### 下载项目
下载到本地的 GOPATH 目录下(后面会提供 go get 的方式以方便使用)，由于这是 Golang 项目，所以需要你本地有 Golang 的运行环境
```
cd ${GOPATH}/src
git clone git@github.com:ItsWewin/go-chat.git
```

### 编译和运行
#### 编译并运行服务端代码
```
go build -o server go-chat/server/main
./server
```
#### 编译并运行客户端代码
```
go build -o client go-chat/client
./client
```

这样就大功告成，你就可以在本地体验本项目了。

## 项目概况
项目目前实现了如下功能
1. 用户注册、登陆
2. 显示所有在线用户列表
3. 发送群消息（目前是发送给在线的所有用户）
4. 私聊某一个用户
5. 按照消息的类型(info, notice, warn, error, success) 使用不同的颜色打印消息（Unix 和 window 均支持）

## 项目效果图
### 注册
![sign-up](https://github.com/ItsWewin/images/raw/master/Chat/sign-up.png)

### 登陆
![sign-in](https://raw.githubusercontent.com/ItsWewin/images/master/Chat/sign-in.png)

### 显示在线用户列表
![online-user-list](https://github.com/ItsWewin/images/raw/master/Chat/online-user-list.png)
### 群聊
![group-message-1.png](https://github.com/ItsWewin/images/raw/master/Chat/group-message-1.png)
![group-message-2.png](https://github.com/ItsWewin/images/raw/master/Chat/group-message-2.png)
### 私聊
![point-to-point.png](https://github.com/ItsWewin/images/raw/master/Chat/point-to-point.png)
![point-to-point2.png](https://github.com/ItsWewin/images/raw/master/Chat/point-to-point2.png)

------------------------------------------------------------------------------------------------------------------------------

# English
## chat room
The chat room developed by Golang, It's a good project to learn Golang and Golang netWork developemt

## usage

### clone this project to your GOPATH

```
cd ${GOPATH}/src
git clone git@github.com:ItsWewin/go-chat.git
```

### build server and run it

```
go build -o server go-chat/server/main
./server
```

### build client and run it
```
go build -o client go-chat/client
./client
```

### Function of this demo
1. User register
2. User login
3. User send group message
4. Show online user list
5. Point-to-point communication
