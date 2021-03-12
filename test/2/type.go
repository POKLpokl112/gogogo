package main

import "fmt"

type myint int

type Book struct {
	name  string
	price myint
}

func (this *Book) pick()  {
	fmt.Println(this.name,"..pick")
}

type MyBook struct{
	Book
	state int
}

func (this *MyBook) pick()  {
	this.Book.pick()
	fmt.Println(this.state)
}

func main() {

	book := Book{name:"qqq",price: 321}
	
	fmt.Println(book)
	book=Book{}
	fmt.Println(book)
	book.name = "123"
	book.price = 222
	fmt.Println(book)
	book.pick()

	myBook := MyBook{state: 1}
	myBook.pick()
	fmt.Println(myBook)
	myBook.Book=book
	myBook.name="test"
	myBook.pick()
	fmt.Println(book)
	fmt.Println(myBook)
	

}
