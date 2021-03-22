package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	q := make(chan *http.Request)
	go Serve(q)
	for i := 0; i < 10; i++ {
		qq := new(http.Request)
		qq.Method = fmt.Sprint(i)
		q <- qq
		fmt.Println("入列,", i) // 等待活动队列清空。
		//time.Sleep(1 * time.Second) // 可能需要很长时间。
	}
	time.Sleep(10 * time.Second)

}

var sem = make(chan int, 5)

func handle(r *http.Request) {
	sem <- 1
	fmt.Println("执行了handler")   // 等待活动队列清空。
	time.Sleep(1 * time.Second) // 可能需要很长时间。
	<-sem                       // 完成；使下一个请求可以运行。
}

func Serve(queue chan *http.Request) {
	for req := range queue {
		sem <- 1
		go func(req *http.Request) {
			fmt.Println("执行了handler", req.Method) // 等待活动队列清空。
			time.Sleep(1 * time.Second)

			<-sem
		}(req)
	}
	// for {
	// 	req := <-queue
	// 	go handle(req) // 无需等待 handle 结束。
	// }
}

func w(r *http.Request) {

}
