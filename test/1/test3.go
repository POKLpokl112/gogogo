package main

import "fmt"

func q(p *int){
	*p=10;
}

func w(p int){
	p=10;
}

func main()  {
	var a int =1;
	fmt.Println(a);
	w(a);
	fmt.Println(a);
	q(&a);
	fmt.Println(a);

}
