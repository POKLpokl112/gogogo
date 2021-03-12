package main

import "fmt"

type Q interface {
	ss()
	qq() string
	ww(w *string)
}

type W struct {
	d *string
}

func (this W) qq() string {
	return *this.d
}

func (this W) ss() {
	fmt.Println("W")
}

func (this W) ww(w *string) {
	this.d = w
}

type R struct {
	d *string
}

func (this R) qq() string {
	return *this.d
}

func (this R) ss() {
	fmt.Println("R")
}

func (this R) ww(w *string) {
	this.d = w
}

func main() {
	var q Q
	var w string
	w = "qwe"
	q = W{&w}

	w = "qqq"
	q.qq()
	q.ss()
	q.ww(&w)
	fmt.Println(q.qq())
	zz(q)

	w = "asd"
	q = R{&w}
	q.qq()
	q.ss()
	w = "rrr"
	q.ww(&w)
	fmt.Println(q.qq())
	zz(q)
}

func zz(s interface{}) {
	var v, ok = s.(string)
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println(s)
}
