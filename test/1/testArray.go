package main

import "fmt"

func main() {
	q := [5]int{1, 2, 3}
	for _, v := range q {
		fmt.Println(v)
	}

	for i := 0; i < len(q); i++ {
		fmt.Println(q[i])
	}

	test(q)
	for _, v := range q {
		fmt.Println(v)
	}

	var w []int = []int{1, 1, 1, 1, 1}
	for _, v := range w {
		fmt.Println(v)
	}
	test1(w)
	for _, v := range w {
		fmt.Println(v)
	}
}

func test(p [5]int) {
	p[3] = 3
}

func test1(p []int) {
	p[3] = 3
}
