package main

import (
	"fmt"
	"day10/demo06/model"
)

func main() {
	// person
	p := model.NewPerson("Tom")
	fmt.Println(*p) //{Tom 0 0}
	p.SetAge(18)
	fmt.Println(*p) //{Tom 18 0}
	p.SetSal(5000)
	fmt.Println(*p) //{Tom 18 5000}
	// account
	a := model.NewAccount("123456")
	fmt.Println(*a)
	a.SetBalance(30)
	fmt.Println(*a)
	a.SetPassword("123456")
	fmt.Println(*a)
}
	