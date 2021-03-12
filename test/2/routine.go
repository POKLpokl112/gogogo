package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func gogogo() {
	i := 1
	for {
		fmt.Println(i)
		i++
		time.Sleep(1 * time.Second)
	}

}

func main() {
	go gogogo()

	go func(j int) bool {

		fmt.Println(j)
		return true
	}(10)

	func() {
		defer fmt.Println("a defer")
		func() {
			defer fmt.Println("b defer")
			fmt.Println("b")
		}()
		fmt.Println("a")
	}()

	i := 1
	for {
		fmt.Println(i)
		i++
		time.Sleep(1 * time.Second)

	}

}
