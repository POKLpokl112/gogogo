package main

import "fmt"

func deferFun() int{
	fmt.Println("defer")
	return 1
	
}

func retrunFun() int{
	fmt.Println("return");
	return 0

}

func test() int {
	defer deferFun()

	return retrunFun()
}

func main(){
	fmt.Println(test())
}