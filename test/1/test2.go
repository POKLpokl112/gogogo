package main

import "fmt"


func Test(a string,b int) (c int,d int)  {
	fmt.Println(a)
	return b,b*2
}

func main()  {
	
	c, d  :=test("1",2);
	fmt.Println(c+d);
	fmt.Println(test("1",2))
}