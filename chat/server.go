package chat

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip        string
	Port      int
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	Message   chan string
}

// NewServer 创建一个server
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// ListenMessage 监听Message广播消息channel的goroutine, 一旦有消息就发送给全部在线user
func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message
		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()
	}
}

// Broadcast 广播消息方法
func (s *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	s.Message <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {
	//当前连接的业务
	fmt.Println("连接建立成功")
	user := NewUser(conn, s)
	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	//接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err:", err)
				return
			}
			//提取用户的消息,删除最后个\n
			msg := string(buf[:n-1])
			//处理消息
			user.DoMessage(msg)
			//用户的任意消息, 代表当前用户是活跃的
			isLive <- true
		}
	}()

	//阻塞
	for {
		select {
		case <-isLive:
			//当前用户活跃的,应该重置定时器
			//不做任何事情, 为了激活select, 更新下面的定时器

		//设置定时器, 10秒之后
		case <-time.After(time.Second * 60):
			//已经超时
			//将当前user客户端强制关闭
			user.SendMsg("你被踢出\n")
			//销毁资源
			close(user.C)
			//关闭连接
			conn.Close()
			//退出当前Handler
			return //runtime.Goexit()
		}
	}
}

// Start 启动服务器
func (s *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
	}

	go s.ListenMessage()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		//do handler
		go s.Handler(conn)
	}
	//close listen socket
	defer listener.Close()

}
