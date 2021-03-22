package main

import (
	"fmt"
	"log"
	"time"
)

type Work struct {
	Id int
}

func main() {
	workChan := make(chan *Work, 10)
	go func() {
		for i := 0; i < 10; i++ {
			workChan <- &Work{i}
		}
	}()

	time.Sleep(1 * time.Second)
	server(workChan)

}

func server(workChan <-chan *Work) {
	for work := range workChan {

		go ggg(work)
	}
}

func ggg(work *Work) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("work failed", work.Id, e, work.Id)
		}
	}()

	g(work)

}

func g(work *Work) {
	if work.Id == 3 {
		panic("来走一个")
	}

	fmt.Println("work work", work.Id)

}
