package main

import (
	"fmt"
	"reflect"
)

type ggg struct {
	name string `info:"qqq" desc:"ddd"`
	id   int    `info:"www"`
}

func print(w interface{}) {
	elem := reflect.TypeOf(w).Elem()

	for i := 0; i < elem.NumField(); i++ {
		info := elem.Field(i).Tag.Get("info")
		desc := elem.Field(i).Tag.Get("desc")

		fmt.Println(info, "---", desc)
	}
}

func main() {
	var a int = 10
	b := reflect.ValueOf(a)
	c := reflect.TypeOf(a)

	fmt.Println(b, "---", c)

	var w ggg
	print(&w)
}
