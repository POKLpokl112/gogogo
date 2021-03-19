package main

import (
	"net"
	"strings"
)

//User ss
type User struct {
	Name   string
	Addr   string
	Conn   net.Conn
	C      chan string
	server *Server
}

//NewUser ss
func NewUser(conn net.Conn, server *Server) *User {
	addr := conn.RemoteAddr().String()
	user := &User{
		addr,
		addr,
		conn,
		make(chan string),
		server,
	}

	go user.ListenMsg()
	return user
}

func (user *User) OnLine() {
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	user.server.BroadCast(user, "上线了")
}

func (user *User) OffLine() {
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()
	user.server.BroadCast(user, "下线了")
}

func (user *User) ReceiveMsg(msg string) {
	user.Conn.Write([]byte(msg + "\n"))
}

func (user *User) SendMsg(msg string) {
	if msg == "who" {
		for _, u := range user.server.OnlineMap {
			user.ReceiveMsg(u.Name + "---当前在线")
		}
	} else if len(msg) > 6 && msg[:7] == "rename|" {
		rename := msg[7:]
		if _, ok := user.server.OnlineMap[rename]; ok {
			user.ReceiveMsg(rename + "用户名已经存在")
			return
		}
		user.server.mapLock.Lock()
		delete(user.server.OnlineMap, user.Name)
		user.Name = rename
		user.server.OnlineMap[rename] = user
		user.server.mapLock.Unlock()

		user.ReceiveMsg("修改成功")
	} else if len(msg) > 2 && msg[:3] == "to|" {
		ss := strings.Split(msg, "|")
		name := ss[1]
		if name == "" {
			user.ReceiveMsg("消息格式为to|name|msg")
			return
		}

		u := user.server.OnlineMap[name]
		if u == nil {
			user.ReceiveMsg("没有找到用户:" + name)
			return
		}

		if len(ss) < 2 || ss[2] == "" {
			user.ReceiveMsg("没有消息")
			return
		}

		u.ReceiveMsg("[" + user.Addr + "]" + user.Name + ":" + ss[2])

	} else {
		user.server.BroadCast(user, msg)
	}
}

//ListenMsg ss
func (user *User) ListenMsg() {
	for {
		msg := <-user.C
		user.Conn.Write([]byte(msg + "\n"))
	}

}
