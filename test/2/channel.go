package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("b end")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("a---", i)
		}

		close(c)
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		if data,ok:=<-c;ok{
			fmt.Println(data)

		}

	}

	fmt.Println("end")
}
