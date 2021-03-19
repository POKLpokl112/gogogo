package main

import "fmt"

func main()  {
	var slice []int = make([]int,3)

	fmt.Printf("len = %d ,cap=%d, detail = %v \n",len(slice),cap(slice),slice)

	if slice == nil {
		fmt.Println("空")
	} else {
		fmt.Println("非空")
	}

	slice = append(slice,1);

	fmt.Printf("len = %d ,cap=%d, detail = %v \n",len(slice),cap(slice),slice)

	s2 := slice
	fmt.Printf("len = %d ,cap=%d, detail = %v \n",len(s2),cap(s2),s2)

	copy(s2,slice)
	fmt.Printf("len = %d ,cap=%d, detail = %v \n",len(s2),cap(s2),s2)
	s2 = append(s2, 3)
	fmt.Printf("len = %d ,cap=%d, detail = %v \n",len(s2),cap(s2),s2)

	s2[1]=5
	//slice = append(slice,6)
	
	fmt.Printf("len = %d ,cap=%d, detail = %v \n",len(s2),cap(s2),s2)
	fmt.Printf("len = %d ,cap=%d, detail = %v \n",len(slice),cap(slice),slice)




}