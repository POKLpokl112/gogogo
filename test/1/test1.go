package main

import "fmt"

const (
	Q, W = iota + 1, iota + 2
	E, R
)

func main() {
	fmt.Println(Q)
	fmt.Println(W)
	fmt.Println(E)
	fmt.Println(R)
}
