package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct{
	Title string `json:title`
	Price int `json:price`
	Actors []string `json:actors`
}

func main()  {
	movie := Movie{Title:"123",Price:12,Actors:[]string{"qwe","123"}}

	fmt.Println(movie)

	q,err:=json.Marshal(movie)

	if(err!=nil){
		fmt.Println(err)
		return
	}

	fmt.Printf("%s",q)

	var w Movie
	err = json.Unmarshal(q,&w)
	if(err!=nil){
		fmt.Println(err)
		return
	}

	fmt.Println(w)
}