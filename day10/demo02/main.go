package main

import (
	"fmt"
)

type Person struct{
	Name string
}

func (p Person) speak(){
	fmt.Println(p.Name,"是一个好人")
}

func (p Person) jisuan(){
	sum := 0
	for i := 1 ; i <= 1000 ; i++{
		sum += i
	}
	fmt.Println(sum)
}

func (p Person) jisuan2(n int){
	sum := 0
	for i := 1 ; i <= n ; i++{
		sum += i
	}
	fmt.Println(sum)
}

func (p Person) getSum(){
	fmt.Println(1+2)
}

func main() {
	var p Person
	p.Name = "Tom"
	p.speak()
	p.jisuan()
	p.jisuan2(10)
	p.getSum()
}





