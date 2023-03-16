package main

import (
	"fmt"
)

type Person struct{
	Name string
}

//给 A 类型绑定一个方法
func (p Person) test() {
	fmt.Println("test()",p.Name)
}

func main() {

	var p Person
	p.Name = "tom"
	p.test()
}



