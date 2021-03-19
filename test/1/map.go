package main

import "fmt"

func main()  {
	q :=make(map[string]string,3)
	q["1"] = "1"
	q["2"] = "2"
	q["3"] = "3"
	fmt.Println(q)
	q["4"]="4"
	fmt.Println(q)

	w := make(map[int]int)
	w[1]=1
	w[2]=2
	fmt.Println(w)

	e := map[int]int{
		1:1,
		2:2,
	}

	fmt.Println(e)
}