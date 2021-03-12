package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp string
	Port     int
	Name     string
	Flag     int
	Conn     net.Conn
}

func NewClient(serverIp string, port int) *Client {

	client := Client{
		ServerIp: serverIp,
		Port:     port,
		Flag:     999,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, port))
	if err != nil {
		fmt.Println("客户端连接出错了，error:", err)
		return nil
	}

	client.Conn = conn

	return &client
}

func (client *Client) publicChat() {
	fmt.Println("输入exit退出")
	var msg string
	fmt.Scanln(&msg)

	for msg != "exit" {
		if len(msg) != 0 {

			_, err := client.Conn.Write([]byte(msg + "\n"))
			if err != nil {
				fmt.Println("send err", err)
				break
			}
		}

		msg = ""
		fmt.Scanln(&msg)
	}
}

func (client *Client) who() {
	if _, err := client.Conn.Write([]byte("who\n")); err != nil {
		fmt.Println("who err ", err)
	}
}

func (client *Client) privateChat() {
	client.who()
	fmt.Println("please input the name to chat,input exit to left")

	var toName string
	fmt.Scanln(&toName)

	for toName != "exit" {
		fmt.Println("please input msg to send,input exit to left")

		var msg string
		fmt.Scanln(&msg)

		for msg != "exit" {
			client.Conn.Write([]byte("to|" + toName + "|" + msg + "\n"))

			msg = ""
			fmt.Scanln(&msg)
		}

		fmt.Println("please input the name to chat,input exit to left")
		toName = ""
		fmt.Scanln(&toName)
	}

}

func (client *Client) rename() {
	fmt.Println("please enter your name")
	var name string
	fmt.Scanln(&name)
	client.Name = name

	if _, err := client.Conn.Write([]byte("rename|" + name + "\n")); err != nil {
		fmt.Println("rename err", err)
		return
	}

}

func (client *Client) Run() {
	for client.Flag != 0 {

		for !client.menu() {

		}

		switch client.Flag {
		case 1:
			fmt.Println("public mode selected")
			client.publicChat()
			break
		case 2:
			fmt.Println("private mode selected")
			client.privateChat()
			break
		case 3:
			fmt.Println("rename selected")
			client.rename()
			break
		}
	}
}

func (client *Client) DoResp() {
	io.Copy(os.Stdout, client.Conn)
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.public mode")
	fmt.Println("2.private mode")
	fmt.Println("3.rename")
	fmt.Println("0.quit")

	if _, err := fmt.Scanln(&flag); err != nil {
		fmt.Println(">>>>>>>>请输入提示数字", err)
		return false
	}

	if flag >= 0 && flag <= 3 {
		client.Flag = flag
		return true
	}
	fmt.Println(">>>>>>>>请输入提示数字")
	return false

}

var serverIp string
var port int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置ip地址(默认127.0.0.1)")
	flag.IntVar(&port, "port", 9999, "端口默认9999")
}
func main() {
	flag.Parse()

	client := NewClient(serverIp, port)
	if client == nil {
		fmt.Println(">>>>>>>>>connect error")
		return
	}
	fmt.Println(">>>>>>>>>>>connected success")

	go client.DoResp()

	client.Run()

}
