package chat

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// NewUser 创建一个用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	//启动监听当前channel的goroutine
	go user.ListenMessage()
	return user
}

func (u *User) Online() {
	//用户上线, 将用户加入到onlineMap中
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()
	//广播当前用户上线消息
	u.server.Broadcast(u, "已上线")
}

func (u *User) Offline() {
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()
	u.server.Broadcast(u, "下线")
}

//给当前user对应的客户端发送消息
func (u *User) SendMsg(msg string) {
	u.conn.Write([]byte(msg))
}

func (u *User) DoMessage(msg string) {
	//查询当前用户列表
	if msg == "who" {
		u.server.mapLock.Lock()
		for _, user := range u.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			u.SendMsg(onlineMsg)
		}
		u.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式: rename|张三
		newName := strings.Split(msg, "|")[1]
		//判断newName是否已使用
		if _, ok := u.server.OnlineMap[newName]; ok {
			u.SendMsg("当前用户名被使用\n")
		} else {
			u.server.mapLock.Lock()
			delete(u.server.OnlineMap, u.Name)
			u.server.OnlineMap[newName] = u
			u.server.mapLock.Unlock()
			u.Name = newName
			u.SendMsg("更新用户名成功:" + u.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		//消息格式: to|张三|消息内容

		//获取对方的用户名
		toUserName := strings.Split(msg, "|")[1]
		if toUserName == "" {
			u.SendMsg("消息格式不正确, 请使用\"to|张三|消息内容\"格式\n")
			return
		}
		//根据用户名,得到对方的user对象
		toUser, ok := u.server.OnlineMap[toUserName]
		if !ok {
			u.SendMsg("未找到在线" + toUserName + "用户\n")
			return
		}
		//获取消息内容,通过对方的user对象发送消息内容
		content := strings.Split(msg, "|")[2]
		if content == "" {
			u.SendMsg("没有消息内容, 请重发\n")
			return
		}
		toUser.SendMsg(u.Name + "对你说:" + content + "\n")
	} else {
		u.server.Broadcast(u, msg)
	}
}

// 监听当前user channel方法, 一旦有消息, 就发送给客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
