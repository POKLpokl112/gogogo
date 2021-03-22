package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fi, err := os.Stat("C:/Users/C5310723")
	if fi.IsDir() {
		fmt.Printf("%s %s is a directory\n", "asd", err)

	}
	fmt.Println("qweqweqwe")
	//time.Sleep(1 * time.Second)
}

type HandlerFuc func(http.ResponseWriter, *http.Request)

func (f HandlerFuc) ServerHttp(w http.ResponseWriter, req *http.Request) {
	f(w, req)
}
