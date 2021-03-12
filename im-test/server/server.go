package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

//Server ..
type Server struct {
	Ip        string
	Port      int
	OnlineMap map[string]*User
	C         chan string
	mapLock   sync.RWMutex
}

//NewServer ...
func NewServer(ip string, port int) *Server {
	return &Server{Ip: ip, Port: port, OnlineMap: make(map[string]*User), C: make(chan string)}

}

//Handler ..
func (server *Server) Handler(con net.Conn) {
	user := NewUser(con, server)

	user.OnLine()

	isAlive := make(chan bool)
	go func() {
		buff := make([]byte, 2048)
		for {
			n, err := con.Read(buff)

			if n == 0 {
				user.OffLine()
				return
			}

			if err != nil {
				fmt.Println("error:", err)

				if err != io.EOF {
					return
				}

			}

			s := string(buff[0 : n-1])
			user.SendMsg(s)
			isAlive <- true
		}
	}()

	fmt.Println(user.Name + "连接成功了")
	//当前阻塞
	for {
		select {
		case <-isAlive:

		case <-time.After(1000 * time.Second):
			user.ReceiveMsg("你被提了")
			close(user.C)
			con.Close()
			return

		}

	}

}

//
func (server *Server) BroadCast(user *User, msg string) {
	server.C <- user.Name + "---" + msg
}

//
func (server *Server) ListenMsg() {
	for {

		msg := <-server.C
		server.mapLock.Lock()
		for _, user := range server.OnlineMap {
			user.C <- msg
		}
		server.mapLock.Unlock()
	}
}

//123sad
func (server *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Printf("err:", err)
		return
	}

	defer listener.Close()

	go server.ListenMsg()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err:", err)
			continue
		}

		go server.Handler(conn)

	}
}
