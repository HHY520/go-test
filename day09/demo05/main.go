package main

import (
	"fmt"
)
type Cat struct{
	Name string
	Age int
	Color string
	Hobby string
} 

func main (){
	var cat1 Cat
	fmt.Println(cat1)
	cat1 = Cat{Name:"haha",Age:5,Color:"红色",Hobby:"吃<・)))><<"}
	fmt.Println(cat1)
	var cat2 *Cat = new(Cat)
	cat2 = &cat1
	fmt.Println(*cat2)
}

